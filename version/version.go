package version

import (
	"fmt"
	"os"
)

type Version struct {
	BotVersion   string
	BotChangelog string
}

func ReadVersion() (*Version, error) {
	var v Version

	// Getting bot version
	version, err := os.ReadFile("version/version.txt")
	if err != nil {
		fmt.Println("[ERROR] Version: impossible to read version.txt file")
		return nil, err
	}
	v.BotVersion = string(version)

	// Getting bot changelog
	changelog, err := os.ReadFile("version/changelog.txt")
	if err != nil {
		fmt.Println("[ERROR] Changelog: impossible to read changelog.txt file")
		return nil, err
	}
	v.BotChangelog = string(changelog)

	return &v, nil
}
