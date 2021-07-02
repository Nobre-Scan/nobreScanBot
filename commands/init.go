package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Green-Tortoises/nobreScanBot/config"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

const COMMAND_STRINGS_PATH = "commands/command_strings.yaml"

var cText *commands

// Getting strings from Yaml file
func init() {
	cText = &commands{}

	if _, e := os.Stat(COMMAND_STRINGS_PATH); e == nil {
		yamlFile, err := ioutil.ReadFile(COMMAND_STRINGS_PATH)
		if err != nil {
			fmt.Println("[ERROR] Commands: impossible to read YAML file")
			os.Exit(1)
		}

		err = yaml.Unmarshal(yamlFile, &cText)
		if err != nil {
			fmt.Println("[ERROR] Commands: impossible to read YAML file")
			os.Exit(1)
		}

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
	switch c[0] {
	case "ajuda":
		sendHelp(s, m)

	case "ping":
		s.ChannelMessageSend(m.ChannelID, cText.Ping)

	default:
		return
	}
}

func RunAsAdmin(s *discordgo.Session, m *discordgo.MessageCreate, admin string, bot *config.Config) {

}
