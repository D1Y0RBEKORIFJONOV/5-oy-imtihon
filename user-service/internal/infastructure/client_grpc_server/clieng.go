package clientgrpcserver

import (
	"fmt"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"user_service_smart_home/internal/config"
)

type ServiceClient interface {
	UserServiceClient() user1.UserServiceClient
	NotificationServiceClient() notificationpb.NotificationServiceClient
	Close() error
}

type serviceClient struct {
	connection                []*grpc.ClientConn
	userService               user1.UserServiceClient
	notificationServiceClient notificationpb.NotificationServiceClient
}

func NewService(cfg *config.Config) (ServiceClient, error) {
	connSoldiersService, err := grpc.NewClient(fmt.Sprintf("%s", cfg.RPCPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	connnotificationServiceClient, err := grpc.NewClient(fmt.Sprintf("%s", cfg.NotificationUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &serviceClient{
		userService:               user1.NewUserServiceClient(connSoldiersService),
		notificationServiceClient: notificationpb.NewNotificationServiceClient(connnotificationServiceClient),
		connection:                []*grpc.ClientConn{connSoldiersService, connnotificationServiceClient},
	}, nil
}

func (s *serviceClient) UserServiceClient() user1.UserServiceClient {
	return s.userService
}
func (s *serviceClient) NotificationServiceClient() notificationpb.NotificationServiceClient {
	return s.notificationServiceClient
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
