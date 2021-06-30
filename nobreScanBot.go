package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Green-Tortoises/nobreScanBot/config"
	"github.com/bwmarrin/discordgo"
)

var BotPrefix = ""

func main() {
	// Reading config file from disk
	botConfig := config.ReadConfig()

	var discord *discordgo.Session
	var err error
	if discord, err = discordgo.New("Bot " + botConfig.Token); err != nil {
		log.Fatal("Error starting bot: ", err)
	}

	BotPrefix = botConfig.BotPrefix

	discord.AddHandler(ready)
	discord.AddHandler(message)
	discord.AddHandler(guildCreate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("NobreScanBot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
	fmt.Println("Powering off bot")
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, BotPrefix+"help")
}

func message(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// check if the message has the right prefix
	if strings.HasPrefix(m.Content, BotPrefix) {
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
