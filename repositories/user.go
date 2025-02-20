package repositories

import (
	"my_super_project/database"
)

func SaveUser(username, hashedPassword string) error {
	query := `
	INSERT INTO users (username, hashed_password) VALUES ($1, $2)
	`

	_, err := database.Db.Exec(query, username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
