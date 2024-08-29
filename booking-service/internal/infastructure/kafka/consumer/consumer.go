package consumer

import (
	"bytes"
	"context"
	"ekzamen_5/booking-service/internal/config"
	clientgrpcserver "ekzamen_5/booking-service/internal/infastructure/client_grpc_server"
	"encoding/json"
	"fmt"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	bookingpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/booking"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"github.com/twmb/franz-go/pkg/kgo"
	"log"
	"log/slog"
	"time"
)

type Consumer struct {
	consumer *kgo.Client
	user     clientgrpcserver.ServiceClient
	cfg      *config.Config
	logger   *slog.Logger
}

func NewConsumer(cfg *config.Config, logger *slog.Logger) (*Consumer, error) {
	var (
		err      error
		consumer *kgo.Client
	)

	for i := 0; i < 1; i++ {
		consumer, err = kgo.NewClient(
			kgo.SeedBrokers("broker:29092"),
			kgo.ConsumeTopics("QQQQQ"),
		)
		if err != nil {
			logger.Error("err", err.Error())
			time.Sleep(1 * time.Millisecond)
			continue
		}
		break
	}
	if err != nil {
		logger.Error("err", err.Error())
		return nil, fmt.Errorf("failed to create Kafka consumer: %v", err)
	}
	user, err := clientgrpcserver.NewService(cfg)
	if err != nil {
		logger.Error("err", err.Error())
		return nil, fmt.Errorf("failed to create Kafka consumer: %v", err)
	}
	return &Consumer{
		consumer: consumer,
		user:     user,
		cfg:      cfg,
		logger:   logger,
	}, nil
}

func (c *Consumer) Consume() {
	const op = "MessageBroker.Kafka.Consume"
	log := c.logger.With(
		slog.String("event", op))
	ctx := context.Background()

	for false {
		fetches := c.consumer.PollFetches(ctx)

		if errs := fetches.Errors(); len(errs) > 0 {
			for _, err := range errs {
				log.Error("Kafka fetch error", err)
			}
			continue
		}

		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			log.Info("partition", fmt.Sprintf("%d", ftp.Partition))

			for _, record := range ftp.Records {
				log.Info("Partition: %d, Offset: %d, Value: %s\n", record.Partition, record.Offset, string(record.Value))
				switch {
				case bytes.Equal(record.Key, c.cfg.MessageBrokerUses.Keys.CreateOrder):
					status, err := c.createOrder(ctx, record.Value)
					if err != nil {
						log.Error("createOrder error", err)
						continue
					}
					log.Info("Create user", status)
				case bytes.Equal(record.Key, c.cfg.MessageBrokerUses.Keys.AddWaitGroup):
					status, err := c.addWaitGroup(ctx, record.Value)
					if err != nil {
						log.Error("addWaitGroup error", err)
						continue
					}
					log.Info("Update user", status)
				default:
					log.Warn("Unknown key", string(record.Key))
				}
			}
		})
	}
}

func (c *Consumer) createOrder(ctx context.Context, value []byte) (bool, error) {
	var req *bookingpb.CreateBookingReq
	err := json.Unmarshal(value, &req)
	if err != nil {
		return false, err
	}
	user, err := c.user.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "id",
		Value: req.UserID,
	})
	if err != nil {
		return false, err
	}
	log.Println("AAAAAAAAAAAAAAAAAAA", user)
	status, err := c.user.BookingServiceClient().CreateBooking(ctx, req)
	if err != nil {
		log.Println("errrr")
		return false, err
	}
	if status != nil && status.IsError {
		_, err = c.user.NotificationServiceClient().AddNotification(ctx, &notificationpb.AddNotificationReq{
			UserId: req.UserID,
			Messages: &notificationpb.CreateMessage{
				SenderName: "BOOKING-SERVICE",
				Status:     "ushbu type dagi honada joy qolamagan!",
			},
		})
		if err != nil {
			return false, err
		}
		_, err = c.user.NotificationServiceClient().SendEmailNotification(ctx, &notificationpb.SendEmailNotificationReq{
			Email:        user.Email,
			SenderName:   "BOOKING-SERVICE",
			Notification: "ushbu type dagi honada joy qolamagan!",
		})
		if err != nil {
			return false, err
		}
		return true, nil
	}

	_, err = c.user.NotificationServiceClient().AddNotification(ctx, &notificationpb.AddNotificationReq{
		UserId: user.Id,
		Messages: &notificationpb.CreateMessage{
			SenderName: "USER-SERVICE",
			Status:     "siz  honaga buyurmani mufaqiyatli yakunladingiz!",
		},
	})
	if err != nil {
		return false, err
	}

	return true, err
}

func (c *Consumer) addWaitGroup(ctx context.Context, value []byte) (bool, error) {
	var req *bookingpb.AddUserToWaitingGroupReq
	err := json.Unmarshal(value, &req)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	user, err := c.user.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "id",
		Value: req.UserId,
	})
	if err != nil {
		return false, err
	}

	_, err = c.user.BookingServiceClient().AddUserToWaitingGroup(ctx, req)
	if err != nil {
		return false, err
	}

	_, err = c.user.NotificationServiceClient().AddNotification(ctx, &notificationpb.AddNotificationReq{
		UserId: user.Id,
		Messages: &notificationpb.CreateMessage{
			SenderName: "BOOKING-SERVICE",
			Status:     fmt.Sprintf("siz %s turidagi honalar uchun  kutush listigda qoshildingiz", req.RoomType),
		},
	})

	_, err = c.user.NotificationServiceClient().SendEmailNotification(ctx, &notificationpb.SendEmailNotificationReq{
		Email:        user.Email,
		SenderName:   "USER-SERVICE",
		Notification: fmt.Sprintf("siz %s turidagi honalar uchun  kutush listigda qoshildingiz", req.RoomType),
	})

	if err != nil {
		c.logger.Error("email:err", err.Error())
	}
	return true, err
}
