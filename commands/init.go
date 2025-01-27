package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/Nobre-Scan/nobreScanBot/config"
	"github.com/Nobre-Scan/nobreScanBot/database"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

const COMMAND_STRINGS_PATH = "commands/command_strings.yaml"

var cText *commands

// Getting strings from Yaml file
func init() {
	cText = &commands{}

	if _, e := os.Stat(COMMAND_STRINGS_PATH); e == nil {
		yamlFile, err := os.ReadFile(COMMAND_STRINGS_PATH)
		if err != nil {
			fmt.Println("[ERROR] Commands: impossible to read YAML file")
			os.Exit(1)
		}

		err = yaml.Unmarshal(yamlFile, &cText)
		if err != nil {
			fmt.Println("[ERROR] Commands: impossible to read YAML file")
			os.Exit(1)
		}

		// Making the caches for the raritys
		calculateRarity()

	} else if os.IsNotExist(e) {
		// Creating YAML file
		yamlConfig, err := yaml.Marshal(exampleFile())
		if err != nil {
			fmt.Println("Error creating config json: ", err)
			os.Exit(1)
		}

		if err = ioutil.WriteFile(COMMAND_STRINGS_PATH, yamlConfig, 0640); err != nil {
			fmt.Println("Error writing file on disk, check if you have the right permissions!", err)
			os.Exit(1)
		}

		fmt.Println("Please make your configuration in " + COMMAND_STRINGS_PATH + ".\nThen restart the app.")
		os.Exit(1)
	}
}

func Run(s *discordgo.Session, m *discordgo.MessageCreate, bot *config.Config) {
	c := strings.Split(m.Content, " ")
	numMangas := func() int64 {
		secondParameter := int64(1)
		if len(c) == 2 {
			var err error
			secondParameter, err = strconv.ParseInt(c[1], 0, 0)
			if err != nil {
				database.LogError(errors.New("[STRCONV] Parse int: " + err.Error()))
			}
		} else {
			if secondParameter > 3 {
				secondParameter = 3
			}
		}
		return secondParameter
	}()

	switch c[0] {
	case "ajuda":
		sendHelp(s, m)

	case "ping":
		sendPing(s, m)

	case "xingaradm":
		sendXingarAdm(s, m, bot)

	case "elogiarstaff":
		sendElogiarStaff(s, m, bot)

	case "mangaaleatorio":
		sendMangaAleatorio(s, m, int(numMangas))

	case "hentaialeatorio":

	case "contar":
		addCountChat(s, m, c[1])

	case "melhorcontagem":
		getCountByGuildId(s, m)

	default:
		return
	}
}

// If someone is boosting the server
func RunAsNitro(s *discordgo.Session, m *discordgo.MessageCreate, bot *config.Config) {

}

// If someone is trying to run a bot command
func RunAsAdmin(s *discordgo.Session, m *discordgo.MessageCreate, admin string, bot *config.Config) {

}
