package mangamodules

import (
	"fmt"
	"strings"

	"github.com/darylhjd/mangodex"
)

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
