package app

import (
	"log/slog"
	grpcapp "user_service_smart_home/internal/app/grpc"
	"user_service_smart_home/internal/config"
	"user_service_smart_home/internal/infastructure/repository/mongodb"
	"user_service_smart_home/internal/services"
	"user_service_smart_home/internal/usecase/user_grpc_service"
	"user_service_smart_home/internal/usecase/user_repository_service"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(cfg config.Config, logger *slog.Logger) *App {
	db, err := mongodb.NewMongoDB(&cfg)
	if err != nil {
		panic(err)
	}
	serviceUseCase := userRepositoryservice.NewUserRepositoryService(db, db, db, db)
	service := services.NewUserService(logger, *serviceUseCase)
	grpcServer := userusecaseservice.NewUserGrpcService(service)

	Server := grpcapp.NewApp(logger, cfg.RPCPort, grpcServer)
	return &App{
		GRPCServer: Server,
	}
}
