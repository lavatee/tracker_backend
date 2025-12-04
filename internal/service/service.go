package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lavatee/tracker_backend/internal/model"
	"github.com/lavatee/tracker_backend/internal/repository"
)

type Users interface {
	SignUp(user model.User) (int, error)
	SignIn(telegramUsername string, password string) (string, string, error)
	GetOneUser(userId int) (model.User, error)
	GetUserReferrals(userId int) ([]model.User, error)
	ParseToken(token string) (jwt.MapClaims, error)
	Refresh(refreshToken string) (string, string, error)
}

type Service struct {
	Users
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repo),
	}
}
