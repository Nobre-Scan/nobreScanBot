package database

func pragmaForeignKeys() {
	DB.Exec("PRAGMA foreign_keys = ON")
}
