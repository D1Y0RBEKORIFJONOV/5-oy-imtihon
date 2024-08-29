package tokens

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user_service_smart_home/internal/config"
	"user_service_smart_home/internal/entity"
)

func NewAccessToken(user *entity.User) (string, error) {
	cfg := config.New()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(cfg.Token.AccessTTL).Unix()
	claims["role"] = "user"

	tokenString, err := token.SignedString([]byte(config.Token()))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewRefreshToken(user *entity.User) (string, error) {
	cfg := config.New()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["exp"] = time.Now().Add(cfg.Token.RefreshTTL).Unix()
	claims["role"] = "user"
	claims["email"] = user.Email

	tokenString, err := token.SignedString([]byte(config.Token()))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateTokens(user *entity.User) (string, string, error) {
	accessToken, err := NewAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := NewRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
