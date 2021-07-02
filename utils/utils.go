package utils

import "github.com/bwmarrin/discordgo"

const NOBRE_SCANBOT_IMAGE = "https://cdn.discordapp.com/avatars/781587364522622997/4474df3f121a4862ae9be173fde14afa.webp"
const NOBRE_COLOR = 16731392
const ADM_FLAG = "<CargoAdm>"

func IsAdmin(s *discordgo.Session, user *discordgo.User, admin string) {
	_, err := s.GuildMembers(admin, user.ID, 0)
	if err != nil {
		// Log error in database
	}

}
