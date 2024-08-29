package consumer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"github.com/twmb/franz-go/pkg/kgo"
	"log"
	"log/slog"
	"time"
	"user_service_smart_home/internal/config"
	clientgrpcserver "user_service_smart_home/internal/infastructure/client_grpc_server"
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
			kgo.SeedBrokers(cfg.MessageBrokerUses.URL),
			kgo.ConsumeTopics(cfg.MessageBrokerUses.Topic),
			kgo.ConsumerGroup("user_service"),
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

	for {
		fetches := c.consumer.PollFetches(ctx)

		if errs := fetches.Errors(); len(errs) > 0 {
			for _, err := range errs {
				log.Error("err", err)
			}
			continue
		}

		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			log.Info("partition", fmt.Sprintf("%d", ftp.Partition))

			for _, record := range ftp.Records {
				log.Info("Partition: %d, Offset: %d, Value: %s\n", record.Partition, record.Offset, string(record.Value))

				switch {
				case bytes.Equal(record.Key, c.cfg.MessageBrokerUses.Keys.Create):
					status, err := c.createUser(ctx, record.Value)
					if err != nil {
						log.Error("err", err.Error())
						continue
					}
					log.Info("Create user", status)
				case bytes.Equal(record.Key, c.cfg.MessageBrokerUses.Keys.Update):
					status, err := c.updateUser(ctx, record.Value)
					if err != nil {
						log.Error("err", err.Error())
						continue
					}
					log.Info("Update user", status)
				case bytes.Equal(record.Key, c.cfg.MessageBrokerUses.Keys.Delete):
					status, err := c.deleteUser(ctx, record.Value)
					if err != nil {
						log.Error("err", err.Error())
						continue
					}
					log.Info("Delete user", status)
				case bytes.Equal(record.Key, c.cfg.MessageBrokerUses.Keys.UpdateEmail):
					status, err := c.updateEmail(ctx, record.Value)
					if err != nil {
						log.Error("err", err.Error())
						continue
					}
					log.Info("Update email", status)
				case bytes.Equal(record.Key, c.cfg.MessageBrokerUses.Keys.UpdatePassword):
					status, err := c.updatePassword(ctx, record.Value)
					if err != nil {
						log.Error("err", err.Error())
						continue
					}
					log.Info("Update password", status)
				}
			}
		})
	}
}

func (c *Consumer) createUser(ctx context.Context, value []byte) (bool, error) {
	var req *user1.CreateUSerReq
	err := json.Unmarshal(value, &req)
	if err != nil {
		return false, err
	}

	status, err := c.user.UserServiceClient().CreateUser(ctx, req)
	if err != nil {
		return false, err
	}

	user, err := c.user.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "email",
		Value: req.Email,
	})
	if err != nil {
		return false, err
	}

	_, err = c.user.NotificationServiceClient().CreateNotification(ctx, &notificationpb.CreateNotificationReq{
		UserId: user.Id,
	})
	if err != nil {
		return false, err
	}
	_, err = c.user.NotificationServiceClient().AddNotification(ctx, &notificationpb.AddNotificationReq{
		UserId: user.Id,
		Messages: &notificationpb.CreateMessage{
			SenderName: "USER-SERVICE",
			Status:     "your are successfully registered",
		},
	})
	if err != nil {
		return false, err
	}

	_, err = c.user.NotificationServiceClient().SendEmailNotification(ctx, &notificationpb.SendEmailNotificationReq{
		Email:        user.Email,
		SenderName:   "USER-SERVICE",
		Notification: "your are successfully registered",
	})

	if err != nil {
		return false, err
	}
	return status.Successfully, err
}

func (c *Consumer) updateUser(ctx context.Context, value []byte) (bool, error) {
	var req *user1.UpdateUserReq
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

	status, err := c.user.UserServiceClient().UpdateUser(ctx, req)
	if err != nil {
		return false, err
	}

	_, err = c.user.NotificationServiceClient().AddNotification(ctx, &notificationpb.AddNotificationReq{
		UserId: user.Id,
		Messages: &notificationpb.CreateMessage{
			SenderName: "USER-SERVICE",
			Status:     "your are successfully updated",
		},
	})

	_, err = c.user.NotificationServiceClient().SendEmailNotification(ctx, &notificationpb.SendEmailNotificationReq{
		Email:        user.Email,
		SenderName:   "USER-SERVICE",
		Notification: "your are successfully updated",
	})

	if err != nil {
		c.logger.Error("email:err", err.Error())
	}
	return status.Successfully, err
}

func (c *Consumer) deleteUser(ctx context.Context, value []byte) (bool, error) {
	var req *user1.DeleteUserReq
	err := json.Unmarshal(value, &req)
	if err != nil {
		return false, err
	}
	user, err := c.user.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "id",
		Value: req.UserId,
	})
	if err != nil {
		return false, err
	}

	status, err := c.user.UserServiceClient().DeleteUser(ctx, req)
	if err != nil {
		return false, err
	}
	_, err = c.user.NotificationServiceClient().SendEmailNotification(ctx, &notificationpb.SendEmailNotificationReq{
		Email:        user.Email,
		SenderName:   "USER-SERVICE",
		Notification: "your are successfully deleted",
	})

	if err != nil {
		c.logger.Error("email:err", err.Error())
	}

	return status.Successfully, err
}

func (c *Consumer) updateEmail(ctx context.Context, value []byte) (bool, error) {
	var req *user1.UpdateEmailReq
	err := json.Unmarshal(value, &req)
	if err != nil {
		return false, err
	}
	user, err := c.user.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "id",
		Value: req.UserId,
	})
	if err != nil {
		return false, err
	}
	status, err := c.user.UserServiceClient().UpdateEmail(ctx, req)
	if err != nil {
		return false, err
	}
	_, err = c.user.NotificationServiceClient().AddNotification(ctx, &notificationpb.AddNotificationReq{
		Messages: &notificationpb.CreateMessage{
			SenderName: "USER-SERVICE",
			Status:     "your are successfully updated email ",
		},
	})
	_, err = c.user.NotificationServiceClient().SendEmailNotification(ctx, &notificationpb.SendEmailNotificationReq{
		Email:        user.Email,
		SenderName:   "USER-SERVICE",
		Notification: "your are successfully updated email ",
	})

	if err != nil {
		c.logger.Error("email:err", err.Error())
	}
	return status.Successfully, err
}
func (c *Consumer) updatePassword(ctx context.Context, value []byte) (bool, error) {
	var req *user1.UpdatePasswordReq
	err := json.Unmarshal(value, &req)
	if err != nil {
		return false, err
	}
	status, err := c.user.UserServiceClient().UpdatePassword(ctx, req)
	if err != nil {
		return false, err
	}
	user, err := c.user.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "id",
		Value: req.UserId,
	})
	if err != nil {
		return false, err
	}

	_, err = c.user.NotificationServiceClient().AddNotification(ctx, &notificationpb.AddNotificationReq{
		UserId: user.Id,
		Messages: &notificationpb.CreateMessage{
			SenderName: "USER-SERVICE",
			Status:     "your are successfully updated  password",
		},
	})
	_, err = c.user.NotificationServiceClient().SendEmailNotification(ctx, &notificationpb.SendEmailNotificationReq{
		Email:        user.Email,
		SenderName:   "USER-SERVICE",
		Notification: "your are successfully updated password",
	})

	if err != nil {
		c.logger.Error("email:err", err.Error())
	}
	return status.Successfully, err
}
