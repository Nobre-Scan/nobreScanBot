package mangamodules

import (
	"github.com/darylhjd/mangodex"
	"gitlab.com/lamados/go-nhentai"
)

var MangadexOn = false
var NhentaiOn = false

var MangadexClient *mangodex.DexClient
var NhentaiClient *nhentai.Client

func Init(mangadexUser string, mangadexPass string) {
	//// MANGADEX
	initMangadex(mangadexUser, mangadexPass)

	//// NHENTAI
	initNhentai()
}
