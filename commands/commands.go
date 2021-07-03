package commands

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/Green-Tortoises/nobreScanBot/config"
	"github.com/Green-Tortoises/nobreScanBot/database"
	"github.com/Green-Tortoises/nobreScanBot/mangamodules"
	"github.com/Green-Tortoises/nobreScanBot/utils"
	"github.com/bwmarrin/discordgo"
)

var elogiarStaffRarity int
var xingarAdmRarity int

// Calculating the raritys
func calculateRarity() {
	elogiarStaffRarity = 0
	xingarAdmRarity = 0

	// Calculating raritys for ElogiarStaff
	for _, r := range cText.Staff.ElogiarStaff {
		elogiarStaffRarity += r.Rarity
	}

	// Calculating raritys for XingarAdm
	for _, x := range cText.Staff.XingarAdm {
		xingarAdmRarity += x.Rarity
	}
}

func sendHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	var embMsg discordgo.MessageEmbed

	embMsg.Title = cText.Ajuda.Title
	embMsg.Description = cText.Ajuda.Description
	embMsg.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: utils.NOBRE_SCANBOT_IMAGE}
	embMsg.Color = utils.NOBRE_COLOR

	// Adding commands to the help command
	embMsg.Fields = make([]*discordgo.MessageEmbedField, len(cText.Ajuda.Commands))

	for i, c := range cText.Ajuda.Commands {
		var field discordgo.MessageEmbedField
		field.Inline = true
		field.Name = c.CommandName
		field.Value = c.Description
		embMsg.Fields[i] = &field
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embMsg)
}

// Send pong
func sendPing(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cText.Ping)
	logPingMessage := fmt.Sprintf("[PING] %s (%s): %s", m.Author.Username, m.Author.ID, cText.Ping)
	if err := database.LogEvent(logPingMessage); err != nil {
		fmt.Println(err)
	}
}

// Send bad words to the Adms
func sendXingarAdm(s *discordgo.Session, m *discordgo.MessageCreate, bot *config.Config) {
	// Choose a random message to send
	choosedValue := rand.Int() % xingarAdmRarity
	var choosedMessage string

	for _, a := range cText.Staff.XingarAdm {
		choosedValue -= a.Rarity

		if choosedValue <= 0 {
			choosedMessage = a.Message
			break
		}
	}

	choosedMessage = replaceFlag2Role(choosedMessage, bot.CargoAdm)
	s.ChannelMessageSend(m.ChannelID, choosedMessage)
}

// Send good words to the Adms
func sendElogiarStaff(s *discordgo.Session, m *discordgo.MessageCreate, bot *config.Config) {
	// Choose a random message to send
	choosedValue := rand.Int() % elogiarStaffRarity
	var choosedMessage string

	for _, a := range cText.Staff.ElogiarStaff {
		choosedValue -= a.Rarity

		if choosedValue <= 0 {
			choosedMessage = a.Message
			break
		}
	}

	choosedMessage = replaceFlag2Role(choosedMessage, bot.CargoAdm)
	s.ChannelMessageSend(m.ChannelID, choosedMessage)
}

// Send a random Manga
func sendMangaAleatorio(s *discordgo.Session, m *discordgo.MessageCreate, numMangas int) {
	// Checking if mangadex is on
	if !mangamodules.MangadexOn {
		s.ChannelMessageSend(m.ChannelID, "O modulo de mangas está desativado!")
		return
	}

	for i := 0; i < numMangas; i++ {
		go func(s *discordgo.Session, m *discordgo.MessageCreate) {
			s.ChannelMessageSendEmbed(m.ChannelID, mangamodules.SendEmbedManga())
		}(s, m)
	}
}

// Replace flag CargoAdm for the real adm role
func replaceFlag2Role(message string, adm_id string) string {
	return strings.ReplaceAll(message, utils.ADM_FLAG, "<@&"+adm_id+">")
}

// Adding a server to the counting command
func addCountChat(s *discordgo.Session, m *discordgo.MessageCreate, channel_id string) {
	if database.CreateCountingData(m.GuildID, channel_id) != nil {
		s.ChannelMessageSend(m.ChannelID, "Erro ao adicionar servidor ao sistema de contagem!")
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("O canal <#%s> agora será usado para contagem!", channel_id))
}

// Counting
func Count(s *discordgo.Session, m *discordgo.MessageCreate) {
	num, err := strconv.ParseInt(m.Content, 0, 0)
	if err != nil {
		return
	}

	data, err := database.GetCountingDataByGuildId(m.GuildID)
	if err != nil {
		return
	}

	if data.ChannelId != m.ChannelID {
		return
	}

	if num == int64(data.CurrentCount)+1 {
		database.UpdateCountingDataByGuild(int(num), m.GuildID)

		if data.BestHighscore < int(num) {
			database.UpdateBestCountingDataByGuild(int(num), m.GuildID)
		}
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")

	} else {
		database.UpdateCountingDataByGuild(0, m.GuildID)
		database.UpdateBestMistakeDataByGuild(m.Author.ID, m.GuildID)
		s.MessageReactionAdd(m.ChannelID, m.ID, "❌")

	}
}

// Getting counting data from a guild
func getCountByGuildId(s *discordgo.Session, m *discordgo.MessageCreate) {
	data, err := database.GetCountingDataByGuildId(m.GuildID)
	if err != nil {
		database.LogError(err)
	}

	best := fmt.Sprintf("O maior número atingido nesse servidor foi de %d!", data.BestHighscore)
	if data.BestMistake != "0" {
		best = fmt.Sprintf("%s\n<@%s> estragou tudo!", best, data.BestMistake)
	}
	s.ChannelMessageSend(m.ChannelID, best)
}
