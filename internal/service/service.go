package service

import (
	"context"
	"mime/multipart"

	"github.com/dgrijalva/jwt-go"
	"github.com/lavatee/tracker_backend/internal/model"
	"github.com/lavatee/tracker_backend/internal/repository"
	"github.com/minio/minio-go/v7"
)

type Users interface {
	SignUp(user model.User) (int, error)
	SignIn(telegramUsername string, password string) (string, string, error)
	GetOneUser(userId int) (model.User, error)
	GetUserReferrals(userId int) ([]model.User, error)
	ParseToken(token string) (jwt.MapClaims, error)
	Refresh(refreshToken string) (string, string, error)
	UpdateUserBalance(adminId int, userId int, coins int, action string) error
}

type Nodes interface {
	GetNextNodes(ctx context.Context, id int64) ([]model.Node, error)
	GetPreviousNodes(ctx context.Context, id int64) ([]model.Node, error)
	UpdateNode(ctx context.Context, id int64, name string, points int, userId int) error
	AddNode(ctx context.Context, parentID int64, name string, points int, userId int) (int64, error)
	GetNodeByID(ctx context.Context, id int64) (model.Node, error)
}

type Achievements interface {
	CreateAchievement(ctx context.Context, ach model.Achievement, fileName string, file multipart.File) (int, error)
	DeleteAchievement(achId int, userId int) error
	GetUserAchievements(userId int) ([]model.Achievement, error)
	GetAchievementsByStatus(status string, userId int) ([]model.Achievement, error)
	ApproveAchievement(achId int, userId int) error
	RejectAchievement(achId int, userId int, comment string) error
	GetAchievementById(achId int) (model.Achievement, error)
}

type Products interface {
	CreateProduct(ctx context.Context, userId int, product model.Product, fileName string, file multipart.File) (int, error)
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
	SetRejectedStatus(orderId int, userId int) error
	SetReadyStatus(orderId int, userId int) error
	SetIssuedStatus(orderid int, userId int) error
	GetUserOrders(userId int) ([]model.Order, error)
}

type Service struct {
	Users
	Nodes
	Achievements
	Products
	Cart
	Orders
}

func NewService(repo *repository.Repository, s3 *minio.Client, bucket string) *Service {
	return &Service{
		Users:        NewUsersService(repo),
		Nodes:        NewNodesService(repo),
		Achievements: NewAchievementsService(repo, s3, bucket),
		Products:     NewProductsService(repo, s3, bucket),
		Cart:         NewCartService(repo),
		Orders:       NewOrdersService(repo),
	}
}
