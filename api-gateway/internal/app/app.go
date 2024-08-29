package app

import (
	htppapp "apigateway/internal/app/htpp"
	"apigateway/internal/config"
	clientgrpcserver "apigateway/internal/infastructure/client_grpc_server"
	redisrepository "apigateway/internal/infastructure/repository/redis"
	bookingservice "apigateway/internal/service/booking"
	hotelservice "apigateway/internal/service/hotel"
	userservice "apigateway/internal/service/user"
	bookingusecase "apigateway/internal/usecase/booking"
	hotelusecase "apigateway/internal/usecase/hotel"
	userusecase "apigateway/internal/usecase/user"
	"log/slog"
)

type App struct {
	HTTPApp *htppapp.App
}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	redisDb := redisrepository.NewRedis(*cfg)
	serviceUser := userusecase.NewUserRepo(redisDb, redisDb, redisDb, redisDb)
	client, err := clientgrpcserver.NewService(cfg)
	if err != nil {
		panic(err)
	}
	userServer := userservice.NewUser(*serviceUser, client, logger)
	user := userusecase.NewUserUseCase(userServer)

	hotelService := hotelservice.NewHotel(logger, client)
	hotelUseCase := hotelusecase.NewHotelServiceUseCase(hotelService)

	bookinService := bookingservice.NewBooking(logger, client)
	bookingUSeCase := bookingusecase.NewBookingUseCase(bookinService)
	server := htppapp.NewApp(logger, cfg.RPCPort, user, hotelUseCase, bookingUSeCase)
	return &App{
		HTTPApp: server,
	}
}
