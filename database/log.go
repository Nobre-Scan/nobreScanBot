package database

import "time"

type Log struct {
	Id       int
	CreateAt int64
	Message  string
}

func createLogDatabase() error {
	query, err := DB.Prepare("CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY, create_at INTEGER, message TEXT)")
	if err != nil {
		return err
	}
	query.Exec()
	return nil
}

// Log important events on database
func LogEvent(eventMesage string) error {
	insert, err := DB.Prepare("INSERT INTO logs (create_at, message) VALUES (?, ?)")
	if err != nil {
		return err
	}
	insert.Exec(time.Now().Unix(), eventMesage)
	return nil
}
