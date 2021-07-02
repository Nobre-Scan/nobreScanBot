package commands

import "fmt"

// Bot commands
type commands struct {
	Ajuda  string `yaml:"Ajuda"`
	Ping   string `yaml:"Ping"`
	Staff  staff  `yaml:"Staff"`
	Mangas mangas `yaml:"Mangas"`
}

// All commands related to the Staff team
type staff struct {
	XingarAdm    []xingarAdm    `yaml:"XingarAdm"`
	ElogiarStaff []elogiarStaff `yaml:"ElogiarStaff"`
}

// Message and its rarity
type xingarAdm struct {
	Message string `yaml:"Mensagem"`
	Rarity  int    `yaml:"Raridade"`
}

// Message and its rarity
type elogiarStaff struct {
	Message string `yaml:"Mensagem"`
	Rarity  int    `yaml:"Raridade"`
}

// All commands related to mangas
type mangas struct {
	MangaAleatorio  string `yaml:"MangaAleatorio"`
	HentaiAleatorio string `yaml:"HentaiAleao"`
}

// Making a simple populated struct
func exampleFile() commands {
	var cText commands
	const EXAMPLE_MESSAGE = "Exemplo de mensagem"
	const EXAMPLE_RARITY = 3

	// Ajuda example
	cText.Ajuda = fmt.Sprintf("%s\n%s", EXAMPLE_MESSAGE, EXAMPLE_MESSAGE)

	// Ping example
	cText.Ping = EXAMPLE_MESSAGE

	// XingarAdm example
	var xingar xingarAdm
	xingar.Message = EXAMPLE_MESSAGE
	xingar.Rarity = EXAMPLE_RARITY
	xingarA := make([]xingarAdm, 0)
	xingarA = append(xingarA, xingar)
	xingarA = append(xingarA, xingar)
	cText.Staff.XingarAdm = xingarA

	// ElogiarStaff example
	var elogiar elogiarStaff
	elogiar.Message = EXAMPLE_MESSAGE
	elogiar.Rarity = EXAMPLE_RARITY
	elogiarStff := make([]elogiarStaff, 0)
	elogiarStff = append(elogiarStff, elogiar)
	elogiarStff = append(elogiarStff, elogiar)
	cText.Staff.ElogiarStaff = elogiarStff

	// Mangas example
	cText.Mangas.MangaAleatorio = EXAMPLE_MESSAGE
	cText.Mangas.HentaiAleatorio = EXAMPLE_MESSAGE

	return cText
}
