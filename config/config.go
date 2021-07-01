package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config path
const ConfigFile = "config.json"

type Config struct {
	Token        string `json:"Token"`
	BotPrefix    string `json:"BotPrefix"`
	CargoAdm     string `json:"CargoAdm"`
	MangadexUser string `json:"MangadexUser"`
	MangadexPass string `json:"MangadexPass"`
	DatabasePath string `json:"DatabasePath"`
}

func ReadConfig() *Config {
	var configData Config

	if _, err := os.Stat(ConfigFile); err == nil {
		// Reading file and extracting values
		var byteFile []byte
		if byteFile, err = ioutil.ReadFile(ConfigFile); err != nil {
			log.Fatal("Error reading config file: ", err)
		}
		json.Unmarshal(byteFile, &configData)

	} else if os.IsNotExist(err) {
		// Creating config file
		var jsonConfig []byte
		if jsonConfig, err = json.MarshalIndent(configData, "", " "); err != nil {
			fmt.Println("Error creating config json: ", err)
		}

		if err = ioutil.WriteFile(ConfigFile, jsonConfig, 0640); err != nil {
			fmt.Println("Error writing file on disk, check if you have the right permissions!", err)
		}

		fmt.Println("Please make your configuration in " + ConfigFile + ".\nThen restart the app.")
		return nil
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

		fmt.Println("Schrodinger: ", err)
		return nil
	}

	return &configData
}
