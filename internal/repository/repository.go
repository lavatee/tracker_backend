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
	UpdateUserBalance(userId int, coins int, action string) error //action типо "+" либо "-"
}

type Nodes interface {
	GetNextNodes(ctx context.Context, id int64) ([]model.Node, error)
	GetPreviousNodes(ctx context.Context, id int64) ([]model.Node, error)
	UpdateNode(ctx context.Context, id int64, name string, points int) error
	AddNode(ctx context.Context, parentID int64, name string, points int) (int64, error)
	GetNodeByID(ctx context.Context, id int64) (model.Node, error)
}

type Achievements interface {
	CreateAchievement(ach model.Achievement) (int, error)
	DeleteAchievement(achId int, userId int) error
	GetUserAchievements(userId int) ([]model.Achievement, error)
	GetPendingAchievements() ([]model.Achievement, error)
	SetAchievementStatus(achId int, status string) error
	GetOneAchievement(achId int) (model.Achievement, error)
}

type Products interface {
	CreateProduct(product model.Product) (int, error)
	GetProducts() ([]model.Product, error)
	GetOneProduct(productId int) (model.Product, error)
	DeleteProduct(productId int) error
	UpdateProduct(product model.Product) error
}

type Cart interface {
	AddProductToCart(productInCart model.ProductInCart) (int, error)
	DeleteProductFromCart(productId int, userId int) error
	CleanUserCart(userId int) error
}

type Orders interface {
	CreateOrder(userId int) (int, error)
	GetOneOrder(orderId int) (model.Order, error)
	GetOrdersByStatus(status string) (model.Order, error)
	GetOrderProducts(orderId int) ([]model.OrderedProduct, error)
	SetOrderStatus(status string) error
	GetUserOrders(userId int) ([]model.Order, error)
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
