package main

import (
	"my_super_project/database"
	"my_super_project/server"
	"my_super_project/utils/logger"
)

func main() {
	logger.InitLoggers()

	err := database.InitDB()
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
	server.Run()
}
