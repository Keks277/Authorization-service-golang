package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/ilyababichev/authorization-service/internal/app/grpc"
	"github.com/ilyababichev/authorization-service/internal/services/auth"
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
	var authService
	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
