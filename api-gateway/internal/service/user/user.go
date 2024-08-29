package userservice

import (
	"apigateway/internal/config"
	"apigateway/internal/entity"
	clientgrpcserver "apigateway/internal/infastructure/client_grpc_server"
	"apigateway/internal/infastructure/producer"
	userusecase "apigateway/internal/usecase/user"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"log/slog"
	"math/rand"
	"time"
)

type User struct {
	user   userusecase.UserRepo
	client clientgrpcserver.ServiceClient
	logger *slog.Logger
	cfg    *config.Config
}

func NewUser(user userusecase.UserRepo,
	client clientgrpcserver.ServiceClient,
	logger *slog.Logger) *User {
	cfg := config.New()
	return &User{
		user:   user,
		client: client,
		logger: logger,
		cfg:    cfg,
	}
}

func (u *User) RegisterUser(ctx context.Context, user entity.CreateUserReq) (entity.StatusMessage, error) {
	const op = "User.RegisterUser"
	log := u.logger.With(
		"method", op)
	users, err := u.client.UserServiceClient().GetAllUser(ctx,
		&user1.GetAllUserReq{
			Filed: "email",
			Value: user.Email,
		})

	if err != nil {
		log.Error("Failed to ger user")
		return entity.StatusMessage{}, errors.New(op + err.Error())
	}
	if users.Count != 0 {
		return entity.StatusMessage{}, errors.New(
			op + " user is already registered")
	}
	log.Info("Checking  password with confirm password")
	if user.Password != user.ConfirmPassword {
		log.Error("invalid  confirm password")
		return entity.StatusMessage{}, errors.New(
			"invalid confirm password")
	}

	rand_int := rand.Int() % 1000
	log.Info("Sending secret code to email")
	_, err = u.client.NotificationServiceClient().SendEmailNotification(ctx, &notificationpb.SendEmailNotificationReq{
		Email:        user.Email,
		SenderName:   "API-GATEWAY",
		Notification: fmt.Sprintf("%d", rand_int),
	})
	err = u.user.SaveUserReq(ctx, entity.UserRegisterReq{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		SecretKey: fmt.Sprintf("%d", rand_int),
	}, 5*time.Minute, "user:register")
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New(
			"failed to save user")
	}
	log.Info("Successfully saved user")

	return entity.StatusMessage{
		Message: "check your email",
	}, nil
}

func (u *User) VerifyUser(ctx context.Context, user entity.VerifyUserReq) (entity.StatusMessage, error) {
	const op = "User.VerifyUser"
	log := u.logger.With(
		"method", op)
	tempUser, err := u.user.GetUserRegister(ctx, user.Email, "user:register")
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New(
			"failed to get user")
	}
	if user.SecretCode != tempUser.SecretKey {
		return entity.StatusMessage{}, errors.New(
			"invalid secret code or email")
	}
	req := user1.CreateUSerReq{
		FirstName: tempUser.FirstName,
		LastName:  tempUser.LastName,
		Email:     tempUser.Email,
		Password:  tempUser.Password,
		Address:   tempUser.Address,
	}
	reqBytes, err := json.Marshal(&req)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to marshal request")
	}
	err = producer.Producer(u.cfg.MessageBrokerUses.Keys.Create, reqBytes, u.cfg.MessageBrokerUses.Topic)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to produce request")
	}

	return entity.StatusMessage{
		Message: "user verified processing",
	}, nil
}

func (u *User) Login(ctx context.Context, user entity.LoginReq) (entity.LoginRes, error) {
	const op = "User.Login"
	log := u.logger.With(
		"method", op)

	log.Info("Called user login")
	tokenres, err := u.client.UserServiceClient().Login(ctx, &user1.LoginReq{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		log.Error(err.Error())
		return entity.LoginRes{}, errors.New(
			op + err.Error())
	}
	log.Info("Successfully logged in")

	return entity.LoginRes{
		Token: entity.Token{
			RefreshToken: tokenres.Tokens.RefreshToken,
			AccessToken:  tokenres.Tokens.AccessToken,
		},
	}, nil
}

func (u *User) UpdateUser(ctx context.Context, user entity.UpdateUserReq) (entity.StatusMessage, error) {
	const op = "User.UpdateUser"
	log := u.logger.With(
		"method", op)

	users, err := u.client.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "id",
		Value: user.UserID,
	})
	if err != nil {
		log.Error(err.Error())
		return entity.StatusMessage{}, errors.New(
			op + err.Error())
	}
	if user.FirstName != "" {
		users.FirstName = user.FirstName
	}
	if user.LastName != "" {
		users.LastName = user.LastName
	}

	req := user1.UpdateUserReq{
		UserId:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	reqBytes, err := json.Marshal(&req)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to marshal request")
	}
	err = producer.Producer(u.cfg.MessageBrokerUses.Keys.Update, reqBytes, u.cfg.MessageBrokerUses.Topic)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to produce request")
	}

	err = u.user.UpdateUser(ctx, &entity.User{
		FirstName: users.FirstName,
		LastName:  users.LastName,
		Email:     users.Email,
		Password:  users.Password,
		ID:        users.Id,
		Profile: entity.Profile{
			FirstName: users.FirstName,
			CreatedAt: users.Profile.CreatedAt,
			UpdatedAt: users.Profile.UpdatedAt,
			DeletedAt: users.Profile.DeletedAt,
		},
	}, "user-get", 1*time.Hour)
	if err != nil {
		return entity.StatusMessage{}, errors.New(
			op + err.Error())
	}
	return entity.StatusMessage{
		Message: "user updated processing",
	}, nil
}

func (u *User) UpdatePassword(ctx context.Context, req entity.UpdatePasswordReq) (entity.StatusMessage, error) {
	const op = "User.UpdatePassword"
	log := u.logger.With(
		"method", op)
	log.Info("Called user update password")

	user, err := u.client.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "id",
		Value: req.UserID,
	})
	if err != nil {
		log.Error(err.Error())
		return entity.StatusMessage{}, errors.New(
			op + err.Error())
	}
	req1 := user1.UpdatePasswordReq{
		UserId:      req.UserID,
		Password:    req.Password,
		NewPassword: req.NewPassword,
	}
	reqBytes, err := json.Marshal(&req1)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to marshal request")
	}
	err = producer.Producer(u.cfg.MessageBrokerUses.Keys.UpdatePassword, reqBytes, u.cfg.MessageBrokerUses.Topic)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to produce request")
	}

	err = u.user.UpdateUser(ctx, &entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  req.NewPassword,
		ID:        user.Id,
		Profile: entity.Profile{
			FirstName: user.FirstName,
			CreatedAt: user.Profile.CreatedAt,
			UpdatedAt: user.Profile.UpdatedAt,
			DeletedAt: user.Profile.DeletedAt,
		},
	}, "user-get", 1*time.Hour)

	if err != nil {
		return entity.StatusMessage{}, errors.New(
			op + err.Error())
	}
	return entity.StatusMessage{
		Message: "user updated processing",
	}, nil
}

func (u *User) UpdateEmail(ctx context.Context, req entity.UpdateEmailReq) (entity.StatusMessage, error) {
	const op = "User.UpdateEmail"
	log := u.logger.With(
		"method", op)

	log.Info("Called user update email")
	user, err := u.client.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "id",
		Value: req.UserID,
	})
	if err != nil {
		log.Error(err.Error())
		return entity.StatusMessage{}, errors.New(
			op + err.Error())
	}

	req1 := user1.UpdateEmailReq{
		UserId:   req.UserID,
		NewEmail: req.NewEmail,
	}
	reqBytes, err := json.Marshal(&req1)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to marshal request")
	}
	err = producer.Producer(u.cfg.MessageBrokerUses.Keys.UpdateEmail, reqBytes, u.cfg.MessageBrokerUses.Topic)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to produce request")
	}

	err = u.user.UpdateUser(ctx, &entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     req.NewEmail,
		Password:  user.Password,
		ID:        user.Id,
		Profile: entity.Profile{
			FirstName: user.FirstName,
			CreatedAt: user.Profile.CreatedAt,
			UpdatedAt: user.Profile.UpdatedAt,
			DeletedAt: user.Profile.DeletedAt,
		},
	}, "user-get", 1*time.Hour)

	if err != nil {
		return entity.StatusMessage{}, errors.New(
			op + err.Error())
	}
	log.Info("Successfully updated email")
	return entity.StatusMessage{
		Message: "user updated processing",
	}, nil
}

func (u *User) DeleteUser(ctx context.Context, req entity.DeleteUserReq) (entity.StatusMessage, error) {
	const op = "User.DeleteUser"
	log := u.logger.With(
		"method", op)

	user, err := u.client.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "id",
		Value: req.UserID,
	})
	if err != nil {
		log.Error(err.Error())
		return entity.StatusMessage{}, errors.New(
			op + err.Error())
	}
	err = u.user.DeleteUser(ctx, "user-get", user.Email)
	if err != nil {
		log.Error(err.Error())
		return entity.StatusMessage{}, errors.New(
			op + err.Error())
	}
	req1 := user1.DeleteUserReq{
		UserId:       req.UserID,
		IsHardDelete: req.IsHardDeleted,
	}
	reqBytes, err := json.Marshal(&req1)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to marshal request")
	}
	err = producer.Producer(u.cfg.MessageBrokerUses.Keys.Delete, reqBytes, u.cfg.MessageBrokerUses.Topic)
	if err != nil {
		log.Info(err.Error())
		return entity.StatusMessage{}, errors.New("failed to produce request")
	}

	return entity.StatusMessage{
		Message: "user deleted processing",
	}, nil
}

func (u *User) saveToCashUser(ctx context.Context, user *entity.User, key string) error {
	err := u.user.UpdateUser(ctx, &entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		ID:        user.ID,
		Profile: entity.Profile{
			FirstName: user.FirstName,
			CreatedAt: user.Profile.CreatedAt,
			UpdatedAt: user.Profile.UpdatedAt,
			DeletedAt: user.Profile.DeletedAt,
			Address:   user.Profile.Address,
		},
	}, key, 1*time.Hour)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUser(ctx context.Context, emails string) (entity.User, error) {
	const op = "User.GetUser"
	log := u.logger.With(
		"method", op)
	log.Info("Called user get user")

	user, err := u.client.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
		Filed: "email",
		Value: emails,
	})
	if err != nil {
		return entity.User{}, errors.New(
			op + err.Error())
	}
	if err := u.saveToCashUser(ctx, &entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		ID:        user.Id,
		Profile: entity.Profile{
			FirstName: user.FirstName,
			CreatedAt: user.Profile.CreatedAt,
			UpdatedAt: user.Profile.UpdatedAt,
			DeletedAt: user.Profile.DeletedAt,
			Address:   user.Profile.Address,
		},
	}, "user-get"); err != nil {
		return entity.User{}, errors.New(
			op + err.Error())
	}

	return entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		ID:        user.Id,
		Profile: entity.Profile{
			FirstName: user.FirstName,
			CreatedAt: user.Profile.CreatedAt,
			UpdatedAt: user.Profile.UpdatedAt,
			DeletedAt: user.Profile.DeletedAt,
			Address:   user.Profile.DeletedAt,
		},
	}, nil
}

func (u *User) GetAllUsers(ctx context.Context, req *entity.GetAllUserReq) (*entity.GetAllUserRes, error) {
	const op = "User.GetAllUsers"
	log := u.logger.With(
		"method", op)
	log.Info("called get all users")
	users, err := u.client.UserServiceClient().GetAllUser(ctx, &user1.GetAllUserReq{
		Filed: req.Field,
		Value: req.Value,
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		log.Error(err.Error())
		return nil, errors.New(
			op + err.Error())
	}
	var resp entity.GetAllUserRes
	resp.Count = users.Count
	for _, user := range users.Users {
		usr := entity.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
			ID:        user.Id,
			Profile: entity.Profile{
				FirstName: user.FirstName,
				CreatedAt: user.Profile.CreatedAt,
				UpdatedAt: user.Profile.UpdatedAt,
				DeletedAt: user.Profile.DeletedAt,
				Address:   user.Profile.Address,
			},
		}
		err = u.user.UpdateUser(ctx, &usr, "user-get", time.Hour*1)
		if err != nil {
			log.Error(err.Error())
			return nil, errors.New(
				op + err.Error())
		}
		resp.Users = append(resp.Users, usr)
	}

	log.Info("called get all users")
	return &resp, nil
}
