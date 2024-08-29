package clientgrpcserver

import (
	"ekzamen_5/booking-service/internal/config"
	"fmt"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	bookingpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/booking"
	hotelpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/hotel"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ServiceClient interface {
	UserServiceClient() user1.UserServiceClient
	NotificationServiceClient() notificationpb.NotificationServiceClient
	HotelServiceClient() hotelpb.HotelServiceClient
	BookingServiceClient() bookingpb.BookingServiceClient
	Close() error
}

type serviceClient struct {
	connection                []*grpc.ClientConn
	userService               user1.UserServiceClient
	hotelService              hotelpb.HotelServiceClient
	bookingServiceClient      bookingpb.BookingServiceClient
	notificationServiceClient notificationpb.NotificationServiceClient
}

func NewService(cfg *config.Config) (ServiceClient, error) {
	connSoldiersService, err := grpc.NewClient(fmt.Sprintf("%s", cfg.UserUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connnotificationServiceClient, err := grpc.NewClient(fmt.Sprintf("%s", cfg.NotificationUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connHotelService, err := grpc.NewClient(fmt.Sprintf("%s", cfg.HotelUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connBookingService, err := grpc.NewClient(fmt.Sprintf("%s", cfg.BookingUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	return &serviceClient{
		userService:               user1.NewUserServiceClient(connSoldiersService),
		notificationServiceClient: notificationpb.NewNotificationServiceClient(connnotificationServiceClient),
		hotelService:              hotelpb.NewHotelServiceClient(connHotelService),
		bookingServiceClient:      bookingpb.NewBookingServiceClient(connBookingService),
		connection:                []*grpc.ClientConn{connSoldiersService, connnotificationServiceClient, connHotelService},
	}, nil
}

func (s *serviceClient) UserServiceClient() user1.UserServiceClient {
	return s.userService
}
func (s *serviceClient) NotificationServiceClient() notificationpb.NotificationServiceClient {
	return s.notificationServiceClient
}
func (s *serviceClient) HotelServiceClient() hotelpb.HotelServiceClient {
	return s.hotelService
}
func (s *serviceClient) BookingServiceClient() bookingpb.BookingServiceClient {
	return s.bookingServiceClient
}
func (s *serviceClient) Close() error {
	var err error
	for _, conn := range s.connection {
		if cer := conn.Close(); cer != nil {
			log.Println("Error while closing gRPC connection:", cer)
			err = cer
		}
	}
	return err
}
