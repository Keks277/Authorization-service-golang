package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/ilyababichev/authorization-service/internal/app/grpc"
	"github.com/ilyababichev/authorization-service/internal/services/auth"
	"github.com/ilyababichev/authorization-service/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	//Инициализация хранилища
	storage, err := postgres.New(storagePath)
	if err != nil {
		panic(err)
	}
	//init auth service
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
