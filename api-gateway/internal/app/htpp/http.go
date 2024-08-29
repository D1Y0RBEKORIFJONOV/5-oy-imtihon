package htppapp

import (
	"apigateway/internal/http/router"
	bookingusecase "apigateway/internal/usecase/booking"
	hotelusecase "apigateway/internal/usecase/hotel"
	userusecase "apigateway/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type App struct {
	Logger *slog.Logger
	Port   string
	Server *gin.Engine
}

func NewApp(logger *slog.Logger, port string, handlerService userusecase.User,
	hotel *hotelusecase.HotelServiceUseCaseImpl, booking *bookingusecase.BookingUseCase) *App {
	sever := router.RegisterRouter(handlerService, hotel, booking)
	return &App{
		Port:   port,
		Server: sever,
		Logger: logger,
	}
}

func (app *App) Start() {
	const op = "app.Start"
	log := app.Logger.With(
		slog.String(op, "Starting server"),
		slog.String("port", app.Port))
	log.Info("Starting server")
	err := app.Server.SetTrustedProxies(nil)
	if err != nil {
		log.Error("Error setting trusted proxies", "error", err)
		return
	}
	err = app.Server.RunTLS(app.Port,
		"/localhost.pem",
		"/localhost-key.pem")
	if err != nil {
		log.Info("Failed to start server", "error", err)
	}
}
