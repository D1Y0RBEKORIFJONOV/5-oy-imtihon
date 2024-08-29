package grpcapp

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	user_grpc_server "user_service_smart_home/internal/grpc/user"
	userusecaseservice "user_service_smart_home/internal/usecase/user_grpc_service"
)

type App struct {
	logger     *slog.Logger
	GrpcServer *grpc.Server
	Port       string
}

func NewApp(log *slog.Logger, port string, usrGrpcServer userusecaseservice.UserGrpcService) *App {
	grpcServer := grpc.NewServer()
	user_grpc_server.RegisterUserServiceServer(grpcServer, usrGrpcServer)
	reflection.Register(grpcServer)
	return &App{
		logger:     log,
		Port:       port,
		GrpcServer: grpcServer,
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
