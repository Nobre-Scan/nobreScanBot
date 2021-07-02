package version

import (
	"fmt"
	"io/ioutil"
)

type Version struct {
	BotVersion   string
	BotChangelog string
}

func ReadVersion() (*Version, error) {
	var v Version

	// Getting bot version
	version, err := ioutil.ReadFile("version/version.txt")
	if err != nil {
		fmt.Println("[ERROR] Version: impossible to read version.txt file")
		return nil, err
	}
	v.BotVersion = string(version)

	// Getting bot changelog
	changelog, err := ioutil.ReadFile("version/changelog.txt")
	if err != nil {
		fmt.Println("[ERROR] Changelog: impossible to read changelog.txt file")
		return nil, err
	}
	v.BotChangelog = string(changelog)

	return &v, nil
}
