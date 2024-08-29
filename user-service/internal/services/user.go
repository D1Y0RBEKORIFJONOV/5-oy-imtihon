package services

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
	"user_service_smart_home/internal/entity"
	"user_service_smart_home/internal/tokens"
	"user_service_smart_home/internal/usecase/user_repository_service"
)

type UserService struct {
	log      *slog.Logger
	userRepo userRepositoryservice.UserRepoUseCase
}

func NewUserService(
	log *slog.Logger,
	userRepo userRepositoryservice.UserRepoUseCase,
) *UserService {
	return &UserService{
		log:      log,
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *entity.CreateUserReq) error {
	const op = "user_service.CreateUser()"
	log := s.log.With(slog.String("method", op))

	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Failed to hash password", err.Error())
		return err
	}
	log.Info("Calling Saver User")
	err = s.userRepo.SaveUser(ctx, &entity.User{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		Email:     req.Email,
		Password:  string(passHash),
		Profile: entity.Profile{
			FirstName: req.FirstName,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Time{},
			Address:   req.Address,
		},
	})
	if err != nil {
		log.Error("Failed to Saver User", err.Error())
		return err
	}

	log.Info("Saver User Success")
	return nil
}

func (u *UserService) Login(ctx context.Context, req *entity.LoginReq) (entity.LoginRes, error) {
	const op = "user_service.Login()"
	log := u.log.With(slog.String("method", op))

	if u.userRepo.IsDeleted(ctx, entity.GetUserReq{Field: "email", Value: req.Email}) {
		log.Error("userRepo.IsDeleted()")
		return entity.LoginRes{}, entity.ErrUserDeleted
	}

	user, err := u.userRepo.GetUser(ctx,
		&entity.GetUserReq{
			Field: "email",
			Value: req.Email,
		})
	if err != nil {
		if errors.Is(err, entity.ErrorConflict) {
			log.Error("Failed to Login", err.Error())
			return entity.LoginRes{}, errors.New("user already exists")
		}
		return entity.LoginRes{}, err
	}
	log.Info("Check User Password")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Error("Failed to Login", err.Error())
		return entity.LoginRes{}, errors.New("invalid password")
	}

	log.Info("Login Success")
	var token entity.Token
	token.AccessToken, token.RefreshToken, err = tokens.GenerateTokens(user)
	if err != nil {
		return entity.LoginRes{}, err
	}
	return entity.LoginRes{
		Token: token,
	}, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *entity.UpdateUserReq) error {
	const op = "user_service.UpdateUser()"
	log := u.log.With(slog.String("method", op))
	log.Info("Calling Saver User")
	if u.userRepo.IsDeleted(ctx, entity.GetUserReq{Field: "id", Value: req.UserID}) {
		log.Error("userRepo.IsDeleted()")
		return entity.ErrUserDeleted
	}
	err := u.userRepo.UpdateUser(ctx, &entity.UpdateUserReq{
		UserID:    req.UserID,
		LastName:  req.LastName,
		FirstName: req.FirstName,
	})
	if err != nil {
		log.Error("Failed to Saver User", err.Error())
		return err
	}
	log.Info("Saver User Success")
	return nil
}

func (u *UserService) UpdateUserPassword(ctx context.Context, req *entity.UpdatePasswordReq) error {
	const op = "user_service.UpdateUserPassword()"
	log := u.log.With(slog.String("method", op))

	if u.userRepo.IsDeleted(ctx, entity.GetUserReq{Field: "id", Value: req.UserID}) {
		log.Error("userRepo.IsDeleted()")
		return entity.ErrUserDeleted
	}
	log.Info("Calling Saver User")
	user, err := u.userRepo.GetUser(ctx, &entity.GetUserReq{
		Field: "id",
		Value: req.UserID,
	})
	if err != nil {
		if errors.Is(err, entity.ErrorNotFound) {
			log.Error("Failed to UpdateUserPassword", err.Error())
			return errors.New("user not found")
		}
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Debug("1", slog.String("hash", user.Password))
		log.Debug("2", slog.String("password", req.Password))
		log.Debug("3", slog.String("password", req.NewPassword))

		log.Error("Failed to Login", err.Error())
		return errors.New("invalid password")
	}

	newPasHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)

	err = u.userRepo.UpdateUserPassword(ctx, string(newPasHash), req.UserID)
	if err != nil {
		log.Error("Failed to UpdateUserPassword", err.Error())
		return err
	}
	log.Info("UpdateUserPassword Success")
	return nil
}

func (u *UserService) UpdateUserEmail(ctx context.Context, req *entity.UpdateEmailReq) error {
	const op = "user_service.UpdateUserEmail()"
	log := u.log.With(slog.String("method", op))

	if u.userRepo.IsDeleted(ctx, entity.GetUserReq{Field: "id", Value: req.UserID}) {
		log.Error("userRepo.IsDeleted()")
		return entity.ErrUserDeleted
	}

	log.Info("Calling Saver User")
	err := u.userRepo.UpdateUserEmail(ctx, &entity.UpdateEmailReq{
		UserID:   req.UserID,
		NewEmail: req.NewEmail,
	})
	if err != nil {
		if errors.Is(err, entity.ErrorNotFound) {
			log.Error("Failed to UpdateUserEmail", err.Error())
			return errors.New("user not found")
		}
		return err
	}
	log.Info("Saver User Success")
	return nil
}

func (u *UserService) GetUser(ctx context.Context, req *entity.GetUserReq) (*entity.User, error) {
	const op = "user_service.GetUser()"
	log := u.log.With(slog.String("method", op))

	if u.userRepo.IsDeleted(ctx, *req) {
		log.Error("userRepo.IsDeleted()")
		return nil, entity.ErrUserDeleted
	}
	log.Info("Calling Saver User")
	user, err := u.userRepo.GetUser(ctx, req)
	if err != nil {
		if errors.Is(err, entity.ErrorNotFound) {
			log.Error("Failed to GetUser", err.Error())
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	log.Info("Saver User Success")
	return user, nil
}

func (u *UserService) GetAllUser(ctx context.Context, req *entity.GetAllUserReq) ([]*entity.User, error) {
	const op = "user_service.GetAllUser()"
	log := u.log.With(slog.String("method", op))

	if u.userRepo.IsDeleted(ctx, entity.GetUserReq{
		Field: req.Field,
		Value: req.Value,
	}) {
		log.Error("userRepo.IsDeleted()")
		return nil, entity.ErrUserDeleted
	}

	log.Info("Calling Saver User")
	users, err := u.userRepo.GetAllUser(ctx, req)
	if err != nil {
		if errors.Is(err, entity.ErrUserIsEmpty) {
			log.Error("Failed to GetAllUser", err.Error())
			return nil, errors.New("user is empty")
		}
		return nil, err
	}
	log.Info("Saver User Success")
	return users, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error {
	const op = "user_service.DeleteUser()"
	log := u.log.With(slog.String("method", op))
	if u.userRepo.IsDeleted(ctx, entity.GetUserReq{Field: "id", Value: req.UserID}) {
		log.Error("userRepo.IsDeleted()")
		return entity.ErrUserDeleted
	}
	log.Info("Calling Saver User")
	err := u.userRepo.DeleteUser(ctx, req)
	if err != nil {
		if errors.Is(err, entity.ErrorNotFound) {
			log.Error("Failed to DeleteUser", err.Error())
			return errors.New("user not found")
		}
		return err
	}
	log.Info("Saver User Success")
	return nil
}
