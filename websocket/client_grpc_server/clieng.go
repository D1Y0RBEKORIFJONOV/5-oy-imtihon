package clientgrpcserver

import (
	"ekzamen_5/websocket/config"
	"fmt"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ServiceClient interface {
	NotificationService() notificationpb.NotificationServiceClient
	Close() error
}

type serviceClient struct {
	connection   []*grpc.ClientConn
	notification notificationpb.NotificationServiceClient
}

func NewService(cfg *config.Config) (ServiceClient, error) {
	connNotification, err := grpc.NewClient(fmt.Sprintf("%s",
		cfg.NotificationURl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &serviceClient{
		notification: notificationpb.NewNotificationServiceClient(connNotification),
		connection:   []*grpc.ClientConn{connNotification},
	}, nil
}

func (s *serviceClient) NotificationService() notificationpb.NotificationServiceClient {
	return s.notification
}
func (s *serviceClient) Close() error {
	var err error
	for _, conn := range s.connection {
		if cerr := conn.Close(); cerr != nil {
			log.Println("Error while closing gRPC connection:", cerr)
			err = cerr
		}
	}
	return err
}
