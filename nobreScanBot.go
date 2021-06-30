package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Green-Tortoises/nobreScanBot/config"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Reading config file from disk
	BotConfig := config.ReadConfig()

	var discord *discordgo.Session
	var err error
	if discord, err = discordgo.New("Bot ", BotConfig.Token); err != nil {
		log.Fatal("Error starting bot: ", err)
	}

	discord.AddHandler(ready)
	discord.AddHandler(message)

	discord.Close()
	fmt.Println("Powering off bot")
}

func ready(s *discordgo.Session, event *discordgo.Ready, botConfig config.Config) {
	s.UpdateGameStatus(0, botConfig.BotPrefix+"help")
}

func message(s *discordgo.Session, m *discordgo.MessageCreate, botConfig config.Config) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// check if the message has the right prefix
	if strings.HasPrefix(m.Content, botConfig.BotPrefix) {
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
			_, _ = s.ChannelMessageSend(channel.ID, "Airhorn is ready! Type !airhorn while in a voice channel to play a sound.")
			return
		}
	}
}
