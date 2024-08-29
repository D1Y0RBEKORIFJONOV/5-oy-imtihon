package app

import (
	grpcapp "ekzamen_5/hotel-service/internal/app/grpc"
	"ekzamen_5/hotel-service/internal/config"
	"ekzamen_5/hotel-service/internal/infastructure/repository/postgresql"
	"ekzamen_5/hotel-service/internal/postgres"
	hotelservices "ekzamen_5/hotel-service/internal/services/hotel"
	hotelusecase "ekzamen_5/hotel-service/internal/usecase/hotel"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	postgresDB, err := postgres.New(cfg)
	if err != nil {
		panic(err)
	}
	db := postgresql.NewSolderRepository(postgresDB, logger)

	serviceUseCase := hotelusecase.NewHotelRepoUseCase(db)

	service := hotelservices.NewHotelService(logger, serviceUseCase)

	serverUseCase := hotelusecase.NewHotelServiceUseCase(service)

	server := grpcapp.NewApp(logger, cfg.RPCPort, serverUseCase)
	return &App{
		GRPCServer: server,
	}
}
