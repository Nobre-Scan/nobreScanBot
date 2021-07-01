package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Init all databases on disk
func Init(dbPath string) error {
	// Initializing database on disk
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return errors.New("[Database] Failed to connect to the database: " + err.Error())
	}

	// Starting basic functions
	pragmaForeignKeys()
	DB.SetMaxOpenConns(1)

	// Adding tables if not exists
	if createLogDatabase() != nil {
		return errors.New("[Error] Log Database: Failed to create")
	}
	if createErrorLogDatabase() != nil {
		return errors.New("[Error] ErrorLog Database: Failed to create")
	}
	if createNobreDatabase() != nil {
		return errors.New("[Error] NobreScan Database: Failed to create")
	}

	fmt.Println("[Databases] All databases created")
	return nil
}

// Close database
func Close() error {
	return DB.Close()
}
