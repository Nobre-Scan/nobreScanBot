package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Green-Tortoises/nobreScanBot/config"
	"github.com/Green-Tortoises/nobreScanBot/mangamodules"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Reading config file from disk
	botConfig := config.ReadConfig()

	discord, err := discordgo.New("Bot " + botConfig.Token)
	if err != nil {
		log.Fatal("Error starting bot: ", err)
	}

	discord.AddHandler(func(s *discordgo.Session, event *discordgo.Ready) { ready(s, event, botConfig.BotPrefix) })
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) { message(s, m, botConfig.BotPrefix) })
	discord.AddHandler(guildCreate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Initialing external modules
	mangamodules.Init(botConfig.MangadexUser, botConfig.MangadexPass)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("NobreScanBot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
	fmt.Println("\nPowering off bot.")
}

func ready(s *discordgo.Session, event *discordgo.Ready, botPrefix string) {
	s.UpdateGameStatus(0, botPrefix+"ajuda")
}

func message(s *discordgo.Session, m *discordgo.MessageCreate, botPrefix string) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// check if the message has the right prefix
	if strings.HasPrefix(m.Content, botPrefix) {
		// run commands
	}
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined.
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "NobreScanBot is ready!")
			return
		}
	}
}
