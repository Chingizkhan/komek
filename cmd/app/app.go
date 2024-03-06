package main

import (
	"komek/config"
	"komek/internal/app"
	"komek/pkg/logger"
	"log"
)

func main() {
	cfg, err := config.New("./config/config.yml")
	if err != nil {
		log.Fatalf("new config: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Debug("debug messages are enabled")
	app.Run(cfg, l)
}
