package database

import (
	"errors"
	"time"
)

type ErrorLog struct {
	Id           int
	CreateAt     int64
	ErrorMessage string
}

func createErrorLogDatabase() error {
	query, err := DB.Prepare("CREATE TABLE IF NOT EXISTS error_logs (id INTEGER PRIMARY KEY, create_at INTEGER, error_message TEXT)")
	if err != nil {
		return err
	}
	query.Exec()
	return nil
}

// Log Error for future reading
func LogError(errMessage error) error {
	insert, err := DB.Prepare("INSERT INTO error_logs (create_at, error_message) VALUES (?, ?)")
	if err != nil {
		return errors.New("[Error] Error Logs: Failed to insert element on database")
	}
	insert.Exec(time.Now().Unix(), errMessage.Error())
	return nil
}
