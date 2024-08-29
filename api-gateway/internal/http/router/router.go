package router

import (
	"apigateway/internal/http/handler"
	"apigateway/internal/http/middleware"
	bookingusecase "apigateway/internal/usecase/booking"
	hotelusecase "apigateway/internal/usecase/hotel"
	userusecase "apigateway/internal/usecase/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(user userusecase.User, hotel *hotelusecase.HotelServiceUseCaseImpl, booking *bookingusecase.BookingUseCase) *gin.Engine {
	userHandler := handler.NewUserServer(user)
	hotelHander := handler.NewHotelService(hotel)
	bookingHandler := handler.NewBooking(booking)
	router := gin.Default()

	router.Use(middleware.Middleware)
	router.Use(middleware.TimingMiddleware)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", userHandler.Login)
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/verify", userHandler.VerifyUser)
		userGroup.PUT("/update", userHandler.UpdateUser)
		userGroup.PUT("/update/password", userHandler.UpdatePassword)
		userGroup.PUT("/update/email", userHandler.UpdateEmail)
		userGroup.DELETE("/delete", userHandler.DeleteUser)
		userGroup.GET("", userHandler.GetUser)
		userGroup.GET("/all", userHandler.GetAllUser)
	}

	hotelGroup := router.Group("/hotel")
	{
		hotelGroup.POST("/create", hotelHander.CreateHotel)
		hotelGroup.GET("/get", hotelHander.GetHotels)
		hotelGroup.PATCH("/update", hotelHander.UpdateHotel)
		hotelGroup.DELETE("/delete/:hotel_id", hotelHander.DeleteHotel)
		hotelGroup.POST("/room/create/:hotel_id", hotelHander.AddRoomToHotels)
		hotelGroup.GET("/room/gets/:hotel_id", hotelHander.GetHotelsRooms)
		hotelGroup.PATCH("/room/update/:hotel_id/:room_id", hotelHander.UpdateHotelsRooms)
		hotelGroup.DELETE("/room/delete/:hotel_id/:room_id", hotelHander.DeleteHotelsRooms)
		hotelGroup.PUT("/room/book/:hotel_id/:room_id", hotelHander.UpdateStatusBook)

	}

	bookingGroup := router.Group("/booking")
	{
		bookingGroup.POST("/create/:hotel_id/:room_type", bookingHandler.CreateBooking)
		bookingGroup.POST("/add/waitlist/:hotel_id/:room_type", bookingHandler.AddUserToWaitList)
	}
	return router
}
