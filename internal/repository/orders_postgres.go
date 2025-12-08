package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lavatee/tracker_backend/internal/model"
)

type OrdersPostgres struct {
	db *sqlx.DB
}

func NewOrdersPostgres(db *sqlx.DB) *OrdersPostgres {
	return &OrdersPostgres{
		db: db,
	}
}

func (r *OrdersPostgres) CreateOrder(userId int) (int, error) {
	type ProductWithPrice struct {
		model.ProductInCart
		Price int `db:"price"`
	}
	var productsInCart []ProductWithPrice
	query := fmt.Sprintf(`
		SELECT pc.*, p.price 
		FROM %s pc 
		JOIN %s p ON pc.product_id = p.id 
		WHERE pc.user_id = $1
	`, productsInCartTable, productsTable)
	if err := r.db.Select(&productsInCart, query, userId); err != nil {
		return 0, err
	}

	if len(productsInCart) == 0 {
		return 0, fmt.Errorf("cart is empty")
	}

	var totalPrice int
	for _, product := range productsInCart {
		totalPrice += product.Price * product.ProductAmount
	}

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var orderId int
	createOrderQuery := fmt.Sprintf("INSERT INTO %s (user_id, status) VALUES ($1, 'pending') RETURNING id", ordersTable)
	if err := tx.QueryRow(createOrderQuery, userId).Scan(&orderId); err != nil {
		return 0, err
	}

	createOrderedProductQuery := fmt.Sprintf(`
		INSERT INTO %s (order_id, product_id, product_amount, price) 
		VALUES ($1, $2, $3, $4)
	`, orderedProductsTable)
	for _, product := range productsInCart {
		if _, err := tx.Exec(createOrderedProductQuery, orderId, product.ProductId, product.ProductAmount, product.Price); err != nil {
			return 0, err
		}
	}

	updateBalanceQuery := fmt.Sprintf("UPDATE %s SET balance = balance - $1 WHERE id = $2", usersTable)
	if _, err := tx.Exec(updateBalanceQuery, totalPrice, userId); err != nil {
		return 0, err
	}

	var balance int
	checkBalanceQuery := fmt.Sprintf("SELECT balance FROM %s WHERE id = $1", usersTable)
	if err := tx.QueryRow(checkBalanceQuery, userId).Scan(&balance); err != nil {
		return 0, err
	}

	if balance < 0 {
		return 0, fmt.Errorf("insufficient balance")
	}

	cleanCartQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", productsInCartTable)
	if _, err := tx.Exec(cleanCartQuery, userId); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return orderId, nil
}

func (r *OrdersPostgres) GetOrderById(orderId int) (model.Order, error) {
	var order model.Order
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", ordersTable)
	if err := r.db.Get(&order, query, orderId); err != nil {
		return model.Order{}, err
	}
	getOrderedProductsQuery := fmt.Sprintf("SELECT op.*, p.name, p.photo_url FROM %s op JOIN %s p ON op.product_id = p.id WHERE op.order_id = $1", orderedProductsTable, productsTable)
	if err := r.db.Get(&order.OrderedProducts, getOrderedProductsQuery, orderId); err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (r *OrdersPostgres) GetOrdersByStatus(status string) ([]model.Order, error) {
	var orders []model.Order
	query := fmt.Sprintf("SELECT * FROM %s WHERE status = $1", ordersTable)
	if err := r.db.Select(&orders, query, status); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrdersPostgres) SetOrderStatus(orderId int, status string) error {
	query := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", ordersTable)
	_, err := r.db.Exec(query, status, orderId)
	return err
}

func (r *OrdersPostgres) GetUserOrders(userId int) ([]model.Order, error) {
	var orders []model.Order
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", ordersTable)
	if err := r.db.Select(&orders, query, userId); err != nil {
		return nil, err
	}
	return orders, nil
}
