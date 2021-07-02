package mangamodules

import (
	"fmt"

	"gitlab.com/lamados/go-nhentai"
)

func initNhentai() {
	NhentaiClient = nhentai.New()
	NhentaiOn = true
	fmt.Println("[Nhentai] Module initialized!")
}
