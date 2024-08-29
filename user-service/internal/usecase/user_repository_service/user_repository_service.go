package userRepositoryservice

import (
	"context"
	"user_service_smart_home/internal/entity"
)

type (
	UserSaver interface {
		SaveUser(ctx context.Context, req *entity.User) error
	}
	ProviderUserRepositoryService interface {
		GetUser(ctx context.Context, req *entity.GetUserReq) (*entity.User, error)
		GetAllUser(ctx context.Context, req *entity.GetAllUserReq) ([]*entity.User, error)
		IsDeleted(ctx context.Context, req entity.GetUserReq) bool
	}
	UpdaterUserRepositoryService interface {
		UpdateUser(ctx context.Context, req *entity.UpdateUserReq) error
		UpdateUserPassword(ctx context.Context, newPassword, userID string) error
		UpdateUserEmail(ctx context.Context, req *entity.UpdateEmailReq) error
	}
	DeleterUserRepositoryService interface {
		DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error
	}
)
type UserRepoUseCase struct {
	saver    UserSaver
	provider ProviderUserRepositoryService
	updater  UpdaterUserRepositoryService
	deleter  DeleterUserRepositoryService
}

func NewUserRepositoryService(saver UserSaver,
	provider ProviderUserRepositoryService,
	Updater UpdaterUserRepositoryService,
	Deleter DeleterUserRepositoryService) *UserRepoUseCase {
	return &UserRepoUseCase{
		saver:    saver,
		provider: provider,
		updater:  Updater,
		deleter:  Deleter,
	}
}

func (u *UserRepoUseCase) IsDeleted(ctx context.Context, req entity.GetUserReq) bool {
	return u.provider.IsDeleted(ctx, req)
}
func (u *UserRepoUseCase) GetUser(ctx context.Context, req *entity.GetUserReq) (*entity.User, error) {
	return u.provider.GetUser(ctx, req)
}
func (u *UserRepoUseCase) GetAllUser(ctx context.Context, req *entity.GetAllUserReq) ([]*entity.User, error) {
	return u.provider.GetAllUser(ctx, req)
}
func (u *UserRepoUseCase) UpdateUser(ctx context.Context, req *entity.UpdateUserReq) error {
	return u.updater.UpdateUser(ctx, req)
}
func (u *UserRepoUseCase) UpdateUserPassword(ctx context.Context, newPassword, userID string) error {
	return u.updater.UpdateUserPassword(ctx, newPassword, userID)
}
func (u *UserRepoUseCase) UpdateUserEmail(ctx context.Context, req *entity.UpdateEmailReq) error {
	return u.updater.UpdateUserEmail(ctx, req)
}
func (u *UserRepoUseCase) DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error {
	return u.deleter.DeleteUser(ctx, req)
}
func (u *UserRepoUseCase) SaveUser(ctx context.Context, req *entity.User) error {
	return u.saver.SaveUser(ctx, req)
}
