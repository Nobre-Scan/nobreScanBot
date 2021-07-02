package commands

import "github.com/bwmarrin/discordgo"

const NOBRE_SCANBOT_IMAGE = "https://cdn.discordapp.com/avatars/781587364522622997/4474df3f121a4862ae9be173fde14afa.webp"
const NOBRE_COLOR = 16731392

func sendHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	var embMsg discordgo.MessageEmbed

	embMsg.Title = cText.Ajuda.Title
	embMsg.Description = cText.Ajuda.Description
	embMsg.Image = &discordgo.MessageEmbedImage{URL: NOBRE_SCANBOT_IMAGE}
	embMsg.Color = NOBRE_COLOR

	s.ChannelMessageSendEmbed(m.ChannelID, &embMsg)
}
