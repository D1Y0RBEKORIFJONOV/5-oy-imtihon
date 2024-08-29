package userusecaseservice

import (
	"context"
	"user_service_smart_home/internal/entity"
)

type UserGrpcService interface {
	CreateUser(ctx context.Context, req *entity.CreateUserReq) error
	Login(ctx context.Context, req *entity.LoginReq) (entity.LoginRes, error)
	UpdateUser(ctx context.Context, req *entity.UpdateUserReq) error
	UpdateUserPassword(ctx context.Context, req *entity.UpdatePasswordReq) error
	UpdateUserEmail(ctx context.Context, req *entity.UpdateEmailReq) error
	GetUser(ctx context.Context, req *entity.GetUserReq) (*entity.User, error)
	GetAllUser(ctx context.Context, req *entity.GetAllUserReq) ([]*entity.User, error)
	DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error
}
type userGrpcService struct {
	user UserGrpcService
}

func NewUserGrpcService(user UserGrpcService) UserGrpcService {
	return &userGrpcService{
		user: user,
	}
}
func (u *userGrpcService) CreateUser(ctx context.Context, req *entity.CreateUserReq) error {
	return u.user.CreateUser(ctx, req)
}
func (u *userGrpcService) Login(ctx context.Context, req *entity.LoginReq) (entity.LoginRes, error) {
	return u.user.Login(ctx, req)
}
func (u *userGrpcService) UpdateUser(ctx context.Context, req *entity.UpdateUserReq) error {
	return u.user.UpdateUser(ctx, req)
}
func (u *userGrpcService) UpdateUserPassword(ctx context.Context, req *entity.UpdatePasswordReq) error {
	return u.user.UpdateUserPassword(ctx, req)
}
func (u *userGrpcService) UpdateUserEmail(ctx context.Context, req *entity.UpdateEmailReq) error {
	return u.user.UpdateUserEmail(ctx, req)
}
func (u *userGrpcService) GetUser(ctx context.Context, req *entity.GetUserReq) (*entity.User, error) {
	return u.user.GetUser(ctx, req)
}
func (u *userGrpcService) GetAllUser(ctx context.Context, req *entity.GetAllUserReq) ([]*entity.User, error) {
	return u.user.GetAllUser(ctx, req)
}
func (u *userGrpcService) DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error {
	return u.user.DeleteUser(ctx, req)
}
