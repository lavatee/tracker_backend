package service

import (
	"context"

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

type Nodes interface {
	GetNextNodes(ctx context.Context, id int64) ([]model.Node, error)
	GetPreviousNodes(ctx context.Context, id int64) ([]model.Node, error)
	UpdateNode(ctx context.Context, id int64, name string, points int, userId int) error
	AddNode(ctx context.Context, parentID int64, name string, points int, userId int) (int64, error)
	GetNodeByID(ctx context.Context, id int64) (model.Node, error)
}

type Service struct {
	Users
	Nodes
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repo),
		Nodes: NewNodesService(repo),
	}
}
