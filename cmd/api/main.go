package main

import "github.com/nekizz/telegram-bot/internal/pkg/server"

// @title Telegram-bot Application
// @description This is a telegram bot management application
// @version 1.0
// @host localhost:8080
// @BasePath /apis/v1

func main() {
	server.Start()
}
