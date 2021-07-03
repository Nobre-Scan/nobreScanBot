package database

type nobre struct {
	Id                    int
	UserId                string
	BestCountingHighscore int
}

// All db for nobrescan
func createNobreDatabase() error {
	query, err := DB.Prepare("CREATE TABLE IF NOT EXISTS nobre (id INTEGER PRIMARY KEY, user_id TEXT, best_counting_highscore INTEGER)")
	if err != nil {
		return err
	}
	query.Exec()
	return nil
}
