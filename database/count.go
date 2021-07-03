package database

import (
	"errors"
	"fmt"
)

type Count struct {
	Id            int
	GuildId       string
	ChannelId     string
	CurrentCount  int
	LastMsgSender string //TODO
	BestHighscore int
	BestMistake   string
}

// Create the counting database
func createCountingDatabase() error {
	query, err := DB.Prepare("CREATE TABLE IF NOT EXISTS contar (id INTEGER PRIMARY KEY, guild_id TEXT NOT NULL UNIQUE, channel_id TEXT, current_count INTEGER, best_highscore INTEGER, best_mistake TEXT)")
	if err != nil {
		return err
	}
	query.Exec()
	return nil
}

// Create counting data for a new guild
func CreateCountingData(guild_id string, channel_id string) error {
	row, err := DB.Prepare("INSERT INTO contar (guild_id, channel_id, current_count, best_highscore, best_mistake) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		errorM := "[ERROR] Counting Database: impossible to insert value"
		fmt.Println(errorM)
		LogError(errors.New(errorM))
		return err
	}

	_, err = row.Exec(guild_id, channel_id, 0, 0, 0)
	if err != nil {
		errorM := "[ERROR] Counting Database: impossible to insert value"
		fmt.Println(errorM)
		LogError(errors.New(errorM))
		return err
	}

	return nil
}

// Update counting
func UpdateCountingDataByGuild(new_count int, guild_id string) error {
	update, err := DB.Prepare("UPDATE contar SET current_count=? WHERE guild_id=?")
	if err != nil {
		errorM := "[ERROR] Counting Database: impossible to update value"
		fmt.Println(errorM)
		LogError(errors.New(errorM))
		return err
	}

	if _, err = update.Exec(new_count, guild_id); err != nil {
		errorM := "[ERROR] Counting Database: impossible to update value"
		fmt.Println(errorM)
		LogError(errors.New(errorM))
		return err
	}

	return nil
}

// Update best counting
func UpdateBestCountingDataByGuild(best_count int, guild_id string) error {
	update, err := DB.Prepare("UPDATE contar SET best_highscore=? WHERE guild_id=?")
	if err != nil {
		errorM := "[ERROR] Counting Database: impossible to update value"
		fmt.Println(errorM)
		LogError(errors.New(errorM))
		return err
	}

	if _, err = update.Exec(best_count, guild_id); err != nil {
		errorM := "[ERROR] Counting Database: impossible to update value"
		fmt.Println(errorM)
		LogError(errors.New(errorM))
		return err
	}

	return nil
}

// Update best mistake
func UpdateBestMistakeDataByGuild(user_id string, guild_id string) error {
	update, err := DB.Prepare("UPDATE contar SET best_mistake=? WHERE guild_id=?")
	if err != nil {
		errorM := "[ERROR] Counting Database: impossible to update value"
		fmt.Println(errorM)
		LogError(errors.New(errorM))
		return err
	}

	if _, err = update.Exec(user_id, guild_id); err != nil {
		errorM := "[ERROR] Counting Database: impossible to update value"
		fmt.Println(errorM)
		LogError(errors.New(errorM))
		return err
	}

	return nil
}

// Get counting data by guild id
func GetCountingDataByGuildId(guild_id string) (*Count, error) {
	rows, err := DB.Query("SELECT * FROM contar WHERE guild_id=?", guild_id)
	if err != nil {
		return nil, err
	}

	var c Count

	for rows.Next() {
		if err := rows.Scan(&c.Id, &c.GuildId, &c.ChannelId, &c.CurrentCount, &c.BestHighscore, &c.BestMistake); err != nil {
			errorM := "[ERROR] Counting Database: impossible to fetch value"
			fmt.Println(errorM)
			LogError(errors.New(errorM))
			return nil, err
		}
	}

	return &c, nil
}
