package repositories

import (
	"database/sql"
	"errors"
	"my_super_project/database"
	"my_super_project/models"
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

func GetUserByUsername(username string) (*models.User, error) {
	query := `
	SELECT username, hashed_password
	FROM users
	WHERE username = $1
	`

	var user models.User
	err := database.Db.QueryRow(query, username).Scan(&user.Username, &user.HashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
