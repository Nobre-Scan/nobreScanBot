package mangamodules

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Green-Tortoises/nobreScanBot/database"
	"github.com/Green-Tortoises/nobreScanBot/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/darylhjd/mangodex"
)

type Manga struct {
	Links       string
	BannerUrl   string
	Title       string
	Description string
}

func initMangadex(mangadexUser string, mangadexPass string) {
	// If no mangadex account provided disable the module
	if strings.Compare(mangadexUser, "") == 0 {
		fmt.Println("[Mangadex] No mangadex account provided! Module disabled!")
		return
	}

	// If account provided try to log in
	c := mangodex.NewDexClient()

	err := c.Login(mangadexUser, mangadexPass)
	if err != nil {
		fmt.Println("[Mangadex Error] Failed to login in magadex service! Module disabled!")
		return
	}

	// Logged in successful
	MangadexClient = c
	MangadexOn = true
	fmt.Println("[Mangadex] Module initialized!")
}

func GetMangaFromMangadex() (*Manga, error) {
	manga, err := MangadexClient.GetRandomManga()
	if err != nil {
		return nil, err
	}

	// Creating manga object
	var mangadexManga Manga

	// Getting the first language
	for _, value := range manga.Data.Attributes.Title {
		mangadexManga.Title = value
		break
	}
	for _, value := range manga.Data.Attributes.Description {
		mangadexManga.Description = value
		break
	}
	for _, value := range manga.Data.Attributes.Links {
		mangadexManga.Links = value
		break
	}

	// Getting cover image
	mangadexCover, _ := MangadexClient.GetMangaCover(manga.Data.ID)
	mangadexManga.BannerUrl = mangadexCover.GetResult()

	return &mangadexManga, nil
}

// Send embedded manga
func SendEmbedManga() *discordgo.MessageEmbed {
	var msgEmb discordgo.MessageEmbed

	for tries := 5; tries > 0; tries-- {
		manga, err := GetMangaFromMangadex()
		if err != nil {
			database.LogError(errors.New("[MANGADEX]: " + err.Error()))
			continue
		}

		// If got manga successful
		// Making the embed message
		msgEmb.Color = utils.NOBRE_COLOR
		msgEmb.Title = manga.Title
		msgEmb.Description = manga.Description
		msgEmb.Image = &discordgo.MessageEmbedImage{URL: manga.BannerUrl}

		tries = 0
	}

	return &msgEmb
}
