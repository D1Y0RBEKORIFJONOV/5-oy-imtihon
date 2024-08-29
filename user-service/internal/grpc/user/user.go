package user_grpc_server

import (
	"context"
	"errors"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	"google.golang.org/grpc"
	"user_service_smart_home/internal/entity"
	"user_service_smart_home/internal/usecase/user_grpc_service"
)

type userServiceServer struct {
	user1.UnimplementedUserServiceServer
	user userusecaseservice.UserGrpcService
}

func RegisterUserServiceServer(GRPCServer *grpc.Server, userService userusecaseservice.UserGrpcService) {
	user1.RegisterUserServiceServer(GRPCServer, &userServiceServer{
		user: userService,
	})
}

func (u *userServiceServer) CreateUser(ctx context.Context, req *user1.CreateUSerReq) (*user1.StatusUser, error) {
	if req.LastName == "" || req.FirstName == "" || req.Password == "" || req.Email == "" {
		return &user1.StatusUser{
				Successfully: false,
			}, entity.ErrBadRequest{
				Err: errors.New("err:bad request"),
			}
	}
	err := u.user.CreateUser(ctx, &entity.CreateUserReq{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		Password:  req.Password,
		Email:     req.Email,
		Address:   req.Address,
	})
	if err != nil {
		return &user1.StatusUser{
			Successfully: false,
		}, err
	}

	return &user1.StatusUser{
		Successfully: true,
	}, nil
}

func (u *userServiceServer) Login(ctx context.Context, req *user1.LoginReq) (*user1.LoginRes, error) {
	if req.Email == "" || req.Password == "" {
		return nil, entity.ErrBadRequest{
			Err: errors.New("err:bad request"),
		}
	}
	tokenRes, err := u.user.Login(ctx, &entity.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &user1.LoginRes{
		Tokens: &user1.Token{
			AccessToken:  tokenRes.Token.AccessToken,
			RefreshToken: tokenRes.Token.RefreshToken,
		},
	}, nil
}

func (u *userServiceServer) UpdateUser(ctx context.Context, req *user1.UpdateUserReq) (*user1.StatusUser, error) {
	if req.LastName == "" && req.FirstName == "" {
		return &user1.StatusUser{
				Successfully: false,
			}, entity.ErrBadRequest{
				Err: errors.New("err:bad request"),
			}
	}
	err := u.user.UpdateUser(ctx, &entity.UpdateUserReq{
		UserID:    req.UserId,
		LastName:  req.LastName,
		FirstName: req.FirstName,
	})
	if err != nil {
		return &user1.StatusUser{
			Successfully: false,
		}, err
	}

	return &user1.StatusUser{
		Successfully: true,
	}, nil
}

func (u *userServiceServer) UpdatePassword(ctx context.Context, req *user1.UpdatePasswordReq) (*user1.StatusUser, error) {
	if req.Password == "" || req.NewPassword == "" {
		return &user1.StatusUser{
				Successfully: false,
			}, entity.ErrBadRequest{
				Err: errors.New("err:bad request"),
			}
	}

	err := u.user.UpdateUserPassword(ctx, &entity.UpdatePasswordReq{
		UserID:      req.UserId,
		Password:    req.Password,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return &user1.StatusUser{
			Successfully: false,
		}, err
	}

	return &user1.StatusUser{
		Successfully: true,
	}, nil
}

func (u *userServiceServer) UpdateEmail(ctx context.Context, req *user1.UpdateEmailReq) (*user1.StatusUser, error) {
	if req.UserId == "" || req.NewEmail == "" {
		return &user1.StatusUser{
				Successfully: false,
			}, entity.ErrBadRequest{
				Err: errors.New("err:bad request"),
			}
	}
	err := u.user.UpdateUserEmail(ctx, &entity.UpdateEmailReq{
		UserID:   req.UserId,
		NewEmail: req.NewEmail,
	})
	if err != nil {
		return &user1.StatusUser{
			Successfully: false,
		}, err
	}

	return &user1.StatusUser{
		Successfully: true,
	}, nil
}

func (u *userServiceServer) GetUser(ctx context.Context, req *user1.GetUserReq) (*user1.User, error) {
	if req.Value == "" || req.Filed == "" {
		return nil, entity.ErrBadRequest{
			Err: errors.New("err:bad request"),
		}
	}
	user, err := u.user.GetUser(ctx, &entity.GetUserReq{
		Value: req.Value,
		Field: req.Filed,
	})
	if err != nil {
		return &user1.User{}, err
	}

	return &user1.User{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Profile: &user1.Profile{
			FirstName: user.Profile.FirstName,
			CreatedAt: user.Profile.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.Profile.UpdatedAt.Format("2006-01-02 15:04:05"),
			DeletedAt: user.Profile.DeletedAt.Format("2006-01-02 15:04:05"),
			Address:   user.Profile.Address,
		},
	}, nil
}

func (u *userServiceServer) GetAllUser(ctx context.Context, req *user1.GetAllUserReq) (*user1.GetAllUserRes, error) {

	users, err := u.user.GetAllUser(ctx, &entity.GetAllUserReq{
		Value: req.Value,
		Field: req.Filed,
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		return &user1.GetAllUserRes{}, err
	}
	var res user1.GetAllUserRes
	for _, user := range users {
		res.Users = append(res.Users, &user1.User{
			Id:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
			Profile: &user1.Profile{
				FirstName: user.Profile.FirstName,
				CreatedAt: user.Profile.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: user.Profile.UpdatedAt.Format("2006-01-02 15:04:05"),
				DeletedAt: user.Profile.DeletedAt.Format("2006-01-02 15:04:05"),
				Address:   user.Profile.Address,
			},
		})
		res.Count += 1
	}
	return &res, nil
}

func (u *userServiceServer) DeleteUser(ctx context.Context, req *user1.DeleteUserReq) (*user1.StatusUser, error) {
	if req.UserId == "" {
		return &user1.StatusUser{
				Successfully: false,
			}, entity.ErrBadRequest{
				Err: errors.New("err:bad request"),
			}
	}

	err := u.user.DeleteUser(ctx, &entity.DeleteUserReq{
		IsHardDelete: req.IsHardDelete,
		UserID:       req.UserId,
	})
	if err != nil {
		return &user1.StatusUser{
			Successfully: false,
		}, err
	}

	return &user1.StatusUser{
		Successfully: true,
	}, nil
}
