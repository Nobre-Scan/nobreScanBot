package database

// All db for nobrescan
func createNobreDatabase() error {
	query, err := DB.Prepare("CREATE TABLE IF NOT EXISTS nobre (id INTEGER PRIMARY KEY)")
	if err != nil {
		return err
	}
	query.Exec()
	return nil
}
