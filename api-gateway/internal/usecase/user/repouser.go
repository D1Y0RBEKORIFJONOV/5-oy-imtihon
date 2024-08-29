package userusecase

import (
	"apigateway/internal/entity"
	"context"
	"time"
)

type (
	Saver interface {
		SaveUserReq(ctx context.Context, user entity.UserRegisterReq, ttl time.Duration, key string) error
	}
	Provider interface {
		GetUserRegister(ctx context.Context, email, key string) (*entity.UserRegisterReq, error)
		GetUser(ctx context.Context, email string, key string) (*entity.User, error)
	}
	Updater interface {
		UpdateUser(ctx context.Context, user *entity.User, key string, ttl time.Duration) error
	}
	Deleter interface {
		DeleteUser(ctx context.Context, key, email string) error
	}
)

type UserRepo struct {
	saver    Saver
	provider Provider
	updater  Updater
	deleter  Deleter
}

func NewUserRepo(saver Saver, provider Provider, updater Updater, delete Deleter) *UserRepo {
	return &UserRepo{
		saver:    saver,
		provider: provider,
		updater:  updater,
		deleter:  delete,
	}
}
func (u *UserRepo) GetUserRegister(ctx context.Context, user, key string) (*entity.UserRegisterReq, error) {
	return u.provider.GetUserRegister(ctx, user, key)
}
func (u *UserRepo) SaveUserReq(ctx context.Context, user entity.UserRegisterReq, ttl time.Duration, key string) error {
	return u.saver.SaveUserReq(ctx, user, ttl, key)
}

func (u *UserRepo) UpdateUser(ctx context.Context, user *entity.User, key string, ttl time.Duration) error {
	return u.updater.UpdateUser(ctx, user, key, ttl)
}

func (u *UserRepo) DeleteUser(ctx context.Context, key, email string) error {
	return u.deleter.DeleteUser(ctx, key, email)
}

func (u *UserRepo) GetUser(ctx context.Context, email string, key string) (*entity.User, error) {
	return u.provider.GetUser(ctx, email, key)
}
