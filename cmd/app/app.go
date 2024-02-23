package main

import (
	"komek/config"
	"komek/internal/app"
	"komek/pkg/logger"
	"log"
	"log/slog"
)

func main() {
	cfg, err := config.New("./config/config.yml")
	if err != nil {
		log.Fatalf("new config: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Info("starting server", slog.String("env", cfg.Log.Level), slog.String("port", cfg.HTTP.Port))
	l.Debug("debug messages are enabled")
	app.Run(cfg, l)
}
