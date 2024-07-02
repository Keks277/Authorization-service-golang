package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ilyababichev/authorization-service/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//Инициализировать объект кофига
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Error("start application", slog.String("env", cfg.Env))

	//Инициализировать логгер

	//инициализировать приложение(app)

	//
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch envLocal {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
