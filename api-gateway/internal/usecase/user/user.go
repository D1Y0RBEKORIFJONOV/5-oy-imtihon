package userusecase

import (
	"apigateway/internal/entity"
	"golang.org/x/net/context"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, user entity.CreateUserReq) (entity.StatusMessage, error)
	VerifyUser(ctx context.Context, user entity.VerifyUserReq) (entity.StatusMessage, error)
	Login(ctx context.Context, user entity.LoginReq) (entity.LoginRes, error)
	UpdateUser(ctx context.Context, user entity.UpdateUserReq) (entity.StatusMessage, error)
	UpdatePassword(ctx context.Context, user entity.UpdatePasswordReq) (entity.StatusMessage, error)
	UpdateEmail(ctx context.Context, user entity.UpdateEmailReq) (entity.StatusMessage, error)
	DeleteUser(ctx context.Context, user entity.DeleteUserReq) (entity.StatusMessage, error)
	GetUser(ctx context.Context, email string) (entity.User, error)
	GetAllUsers(ctx context.Context, req *entity.GetAllUserReq) (*entity.GetAllUserRes, error)
}

type User struct {
	user UserUseCase
}

func NewUserUseCase(user UserUseCase) User {
	return User{user: user}
}

func (u *User) RegisterUser(ctx context.Context, user entity.CreateUserReq) (entity.StatusMessage, error) {
	return u.user.RegisterUser(ctx, user)
}

func (u *User) VerifyUser(ctx context.Context, user entity.VerifyUserReq) (entity.StatusMessage, error) {
	return u.user.VerifyUser(ctx, user)
}

func (u *User) Login(ctx context.Context, user entity.LoginReq) (entity.LoginRes, error) {
	return u.user.Login(ctx, user)
}
func (u *User) UpdateUser(ctx context.Context, user entity.UpdateUserReq) (entity.StatusMessage, error) {
	return u.user.UpdateUser(ctx, user)
}
func (u *User) UpdatePassword(ctx context.Context, user entity.UpdatePasswordReq) (entity.StatusMessage, error) {
	return u.user.UpdatePassword(ctx, user)
}

func (u *User) UpdateEmail(ctx context.Context, user entity.UpdateEmailReq) (entity.StatusMessage, error) {
	return u.user.UpdateEmail(ctx, user)
}
func (u *User) DeleteUser(ctx context.Context, user entity.DeleteUserReq) (entity.StatusMessage, error) {
	return u.user.DeleteUser(ctx, user)
}

func (u *User) GetUser(ctx context.Context, email string) (entity.User, error) {
	return u.user.GetUser(ctx, email)
}

func (u *User) GetAllUsers(ctx context.Context, req *entity.GetAllUserReq) (*entity.GetAllUserRes, error) {
	return u.user.GetAllUsers(ctx, req)

}
