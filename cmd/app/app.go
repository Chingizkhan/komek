package main

import (
	"komek/config"
	"komek/internal/app"
	"komek/pkg/logger"
	"log"
)

// @title API for Komek project
// @version 1
// @description Some description

// @contact.name Chalov Chingizkhan
// @contact.url git@github.com:Chingizkhan/komek.git
// @contact.email cchalovv@mail.ru

// @host localhost:8887
// @BasePath /api/v1
func main() {
	cfg, err := config.New("./config/config.yml")
	if err != nil {
		log.Fatalf("new config: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Debug("debug messages are enabled")
	app.Run(cfg, l)
}
