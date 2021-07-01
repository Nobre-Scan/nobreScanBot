package mangamodules

import (
	"fmt"
	"strings"

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
	mangadexOn = true
	fmt.Println("[Mangadex] Module initialized!")
}

func getMangaFromMangadex() (*Manga, error) {
	manga, err := MangadexClient.GetRandomManga()
	if err != nil {
		return nil, err
	}

	// Creating manga object
	var mangadexManga Manga

	// Getting the first language
	for t := range manga.Data.Attributes.Title {
		mangadexManga.Title = t
		break
	}
	for d := range manga.Data.Attributes.Description {
		mangadexManga.Description = d
		break
	}
	for l := range manga.Data.Attributes.Links {
		mangadexManga.Links = l
		break
	}

	// Getting cover image
	mangadexCover, _ := MangadexClient.GetMangaCover(manga.Data.ID)
	mangadexManga.BannerUrl = mangadexCover.Data.Attributes.FileName

	return &mangadexManga, nil
}
