package grpcapp

import (
	serverhotel "ekzamen_5/hotel-service/internal/grpc/hotel"
	hotelusecase "ekzamen_5/hotel-service/internal/usecase/hotel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type App struct {
	logger     *slog.Logger
	GrpcServer *grpc.Server
	Port       string
}

func NewApp(logger *slog.Logger, port string, hotel *hotelusecase.HotelServiceUseCaseImpl) *App {
	grpcServer := grpc.NewServer()
	serverhotel.RegisterHotelServer(grpcServer, hotel)
	reflection.Register(grpcServer)
	return &App{
		logger:     logger,
		GrpcServer: grpcServer,
		Port:       port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.App.Run"
	log := a.logger.With(
		slog.String("method", op),
		slog.String("port", a.Port))

	l, err := net.Listen("tcp", a.Port)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("starting gRPC server on port", slog.String("port", a.Port))
	err = a.GrpcServer.Serve(l)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}
func (a *App) Stop() {
	log := a.logger.With("port", a.Port)
	log.Info("stopping server")
	a.GrpcServer.GracefulStop()
}
