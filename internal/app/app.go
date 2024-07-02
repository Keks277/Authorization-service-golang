package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/ilyababichev/authorization-service/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenttl time.Duration,
) *App {
	//Инициализация хранилища

	//init auth service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
