package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lavatee/tracker_backend/internal/model"
)

type Users interface {
	SignUp(user model.User) (int, error)
	SignIn(telegramUsername string, passwordHash string) (model.User, error)
	GetOneUser(userId int) (model.User, error)
	GetUserReferrals(userId int) ([]model.User, error)
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUsersPostgres(db),
	}
}
