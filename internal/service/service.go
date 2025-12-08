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

type Achievements interface {
	CreateAchievement(ach model.Achievement) (int, error)
	DeleteAchievement(achId int, userId int) error
	GetUserAchievements(userId int) ([]model.Achievement, error)
	GetAchievementsByStatus(status string, userId int) ([]model.Achievement, error)
	ApproveAchievement(achId int, userId int) error
	RejectAchievement(achId int, userId int) error
	GetAchievementById(achId int) (model.Achievement, error)
}

type Products interface {
	CreateProduct(product model.Product, userId int) (int, error)
	GetProducts() ([]model.Product, error)
	GetProductById(productId int) (model.Product, error)
	DeleteProduct(productId int, userId int) error
	UpdateProduct(product model.Product, userId int) error
}

type Cart interface {
	AddProductToCart(productInCart model.ProductInCart) (int, error)
	UpdateProductInCartAmount(productId int, userId int, amount int) error
	GetUserCart(userId int) ([]model.ProductInCart, error)
	DeleteProductFromCart(productId int, userId int) error
	CleanUserCart(userId int) error
}

type Orders interface {
	CreateOrder(userId int) (int, error)
	GetOrderById(orderId int) (model.Order, error)
	GetOrdersByStatus(status string) ([]model.Order, error)
	SetOrderStatus(orderId int, status string) error
	GetUserOrders(userId int) ([]model.Order, error)
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
