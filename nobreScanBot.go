package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Green-Tortoises/nobreScanBot/commands"
	"github.com/Green-Tortoises/nobreScanBot/config"
	"github.com/Green-Tortoises/nobreScanBot/database"
	"github.com/Green-Tortoises/nobreScanBot/mangamodules"
	"github.com/Green-Tortoises/nobreScanBot/version"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Reading config file from disk
	bot := config.ReadConfig()
	if bot == nil {
		return
	}

	// Reading bot version and changelog
	botVersion, err := version.ReadVersion()
	if err != nil {
		return
	}

	discord, err := discordgo.New("Bot " + bot.Token)
	if err != nil {
		fmt.Println("Error starting bot: ", err)
		return
	}

	discord.AddHandler(func(s *discordgo.Session, event *discordgo.Ready) { ready(s, event, bot.BotPrefix, botVersion) })
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) { message(s, m, bot) })

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	// Initialing external modules
	mangamodules.Init(bot.MangadexUser, bot.MangadexPass)
	if err = database.Init(bot.DatabasePath); err != nil {
		fmt.Println(err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("NobreScanBot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close databases on disk
	database.Close()

	// Cleanly close down the Discord session.
	discord.Close()
	fmt.Println("\nPowering off bot.")
}

func ready(s *discordgo.Session, event *discordgo.Ready, botPrefix string, version *version.Version) {
	gameStatus := fmt.Sprintf("%sajuda - Versão do bot: %s", botPrefix, version.BotVersion)
	s.UpdateGameStatus(0, gameStatus)
}

func message(s *discordgo.Session, m *discordgo.MessageCreate, bot *config.Config) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// check if the message has the right prefix
	if strings.HasPrefix(m.Content, bot.BotPrefix) {
		// Remove prefix from command, Trim and set all the command to lower case
		m.Content = strings.Replace(m.Content, bot.BotPrefix, "", 1)
		m.Content = strings.Trim(m.Content, " ")
		m.Content = strings.ToLower(m.Content)

		// Run commands
		commands.Run(s, m, bot)

	}
}
