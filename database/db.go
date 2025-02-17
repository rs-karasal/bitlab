package database

import (
	"database/sql"
	"fmt"
	"my_super_project/config"
	"my_super_project/utils/logger"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

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
