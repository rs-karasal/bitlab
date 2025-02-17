package database

import (
	"database/sql"
	"my_super_project/utils/logger"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() error {
	dsn := "host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"

	var err error
	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	if err := Db.Ping(); err != nil {
		Db.Close()
		return err
	}

	logger.InfoLog.Print("Connected to database")
	return nil
}
