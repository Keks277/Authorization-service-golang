package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilyababichev/authorization-service/internal/app"
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

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go application.GRPCSrv.MustRun()
	//Инициализировать логгер

	//инициализировать приложение(app)

	//Создаем канал для сигнала от ОС
	stop := make(chan os.Signal, 1)

	//Сообщаем какие сигналы мы ждем
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	//Ждем сигнал в канале
	sign := <-stop

	//Мягко останавливаем сервер
	application.GRPCSrv.Stop()

	log.Info("Server stopped with " + sign.String())
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
