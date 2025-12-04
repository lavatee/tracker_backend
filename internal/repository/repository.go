package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lavatee/tracker_backend/internal/model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Users interface {
	SignUp(user model.User) (int, error)
	SignIn(telegramUsername string, passwordHash string) (model.User, error)
	GetOneUser(userId int) (model.User, error)
	GetUserReferrals(userId int) ([]model.User, error)
	CheckIsAdmin(userId int) bool
}

type Nodes interface {
	GetNextNodes(ctx context.Context, id int64) ([]model.Node, error)
	GetPreviousNodes(ctx context.Context, id int64) ([]model.Node, error)
	UpdateNode(ctx context.Context, id int64, name string, points int) error
	AddNode(ctx context.Context, parentID int64, name string, points int) (int64, error)
	GetNodeByID(ctx context.Context, id int64) (model.Node, error)
}

type Repository struct {
	Users
	Nodes
}

func NewRepository(db *sqlx.DB, neoDriver neo4j.DriverWithContext) *Repository {
	return &Repository{
		Users: NewUsersPostgres(db),
		Nodes: NewNodesNeo(neoDriver),
	}
}
