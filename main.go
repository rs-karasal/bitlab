package main

import (
	"log"
	"my_super_project/config"
	"my_super_project/database"
	"my_super_project/server"
	"my_super_project/utils/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("No .env file found...")
	}

	logger.InitLoggers()

	err = database.InitDB(cfg)
	if err != nil {
		logger.ErrorLog.Print("Can not connect to DB")
		return
	}
	defer func() {
		if database.Db != nil {
			database.Db.Close()
		}
	}()

	// Запускаем наш сервер
	server.Run(cfg)
}
