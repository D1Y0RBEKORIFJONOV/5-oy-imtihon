package app

import (
	grpcapp "ekzamen_5/booking-service/internal/app/grpc"
	"ekzamen_5/booking-service/internal/config"
	clientgrpcserver "ekzamen_5/booking-service/internal/infastructure/client_grpc_server"
	repositorybooking "ekzamen_5/booking-service/internal/infastructure/repository/booking"
	"ekzamen_5/booking-service/internal/postgres"
	bookingservices "ekzamen_5/booking-service/internal/services/booking"
	bookingusecase "ekzamen_5/booking-service/internal/usecase/booking"
	"log/slog"
)

type App struct {
	GrpcServer *grpcapp.App
}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	postgresDb, err := postgres.New(cfg)
	if err != nil {
		panic(err)
	}

	db := repositorybooking.NewSolderRepository(postgresDb, logger)
	serviceUseCase := bookingusecase.NewBookingRepoUseCase(db)

	grpcClient, err := clientgrpcserver.NewService(cfg)
	if err != nil {
		panic(err)
	}
	service := bookingservices.NewBookingService(logger, grpcClient, serviceUseCase)

	serverUseCase := bookingusecase.NewBookingUseCase(service)
	server := grpcapp.NewApp(cfg.RPCPort, logger, serverUseCase)
	return &App{
		GrpcServer: server,
	}
}
