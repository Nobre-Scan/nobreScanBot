package utils

import "github.com/bwmarrin/discordgo"

func IsAdmin(s *discordgo.Session, user *discordgo.User, admin string) {
	_, err := s.GuildMembers(admin, user.ID, 1000)
	if err != nil {
		// Log error in database
	}

}