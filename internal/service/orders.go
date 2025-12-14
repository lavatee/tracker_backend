package service

import (
	"fmt"

	"github.com/lavatee/tracker_backend/internal/model"
	"github.com/lavatee/tracker_backend/internal/repository"
)

const (
	rejectedOrderStatus = "rejected"
	readyOrderStatus    = "ready"
	issuedOrderStatus   = "issued"
)

type OrdersService struct {
	repo *repository.Repository
}

func NewOrdersService(repo *repository.Repository) *OrdersService {
	return &OrdersService{
		repo: repo,
	}
}

func (s *OrdersService) CreateOrder(userId int) (int, error) {
	return s.repo.Orders.CreateOrder(userId)
}

func (s *OrdersService) GetOrderById(orderId int) (model.Order, error) {
	return s.repo.Orders.GetOrderById(orderId)
}

func (s *OrdersService) GetOrdersByStatus(status string) ([]model.Order, error) {
	return s.repo.Orders.GetOrdersByStatus(status)
}

func (s *OrdersService) SetRejectedStatus(orderId int, userId int) error {
	if isAdmin := s.repo.Users.CheckIsAdmin(userId); !isAdmin {
		return fmt.Errorf("user is not an admin")
	}
	return s.repo.Orders.SetOrderStatus(orderId, rejectedOrderStatus)
}

func (s *OrdersService) SetReadyStatus(orderId int, userId int) error {
	if isAdmin := s.repo.Users.CheckIsAdmin(userId); !isAdmin {
		return fmt.Errorf("user is not an admin")
	}
	return s.repo.Orders.SetOrderStatus(orderId, readyOrderStatus)
}

func (s *OrdersService) SetIssuedStatus(orderId int, userId int) error {
	if isAdmin := s.repo.Users.CheckIsAdmin(userId); !isAdmin {
		return fmt.Errorf("user is not an admin")
	}
	return s.repo.Orders.SetOrderStatus(orderId, rejectedOrderStatus)
}

func (s *OrdersService) GetUserOrders(userId int) ([]model.Order, error) {
	return s.repo.Orders.GetUserOrders(userId)
}
