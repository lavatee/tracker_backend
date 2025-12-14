package service

import (
	"github.com/lavatee/tracker_backend/internal/model"
	"github.com/lavatee/tracker_backend/internal/repository"
)

type CartService struct {
	repo *repository.Repository
}

func NewCartService(repo *repository.Repository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) AddProductToCart(productInCart model.ProductInCart) (int, error) {
	return s.repo.Cart.AddProductToCart(productInCart)
}

func (s *CartService) DeleteProductFromCart(productId int, userId int) error {
	return s.repo.Cart.DeleteProductFromCart(productId, userId)
}

func (s *CartService) GetUserCart(userId int) ([]model.ProductInCart, error) {
	return s.repo.Cart.GetUserCart(userId)
}

func (s *CartService) UpdateProductInCartAmount(productId int, userId int, amount int) error {
	return s.repo.Cart.UpdateProductInCartAmount(productId, userId, amount)
}

func (s *CartService) CleanUserCart(userId int) error {
	return s.repo.Cart.CleanUserCart(userId)
}
