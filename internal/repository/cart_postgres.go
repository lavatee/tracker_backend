package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lavatee/tracker_backend/internal/model"
)

type CartPostgres struct {
	db *sqlx.DB
}

func NewCartPostgres(db *sqlx.DB) *CartPostgres {
	return &CartPostgres{
		db: db,
	}
}

func (r *CartPostgres) AddProductToCart(productInCart model.ProductInCart) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, product_id, product_amount) VALUES ($1, $2, $3) RETURNING id", productsInCartTable)
	row := r.db.QueryRow(query, productInCart.UserId, productInCart.ProductId, productInCart.ProductAmount)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CartPostgres) UpdateProductInCartAmount(productId int, userId int, amount int) error {
	query := fmt.Sprintf("UPDATE %s SET product_amount = $1 WHERE product_id = $1", productsInCartTable)
	_, err := r.db.Exec(query, amount, productId)
	return err
}

func (r *CartPostgres) GetUserCart(userId int) ([]model.ProductInCart, error) {
	var productsInCart []model.ProductInCart
	query := fmt.Sprintf("SELECT pc.*, p.name, p.photo_url FROM %s pc JOIN %s p ON pc.product_id = p.id WHERE pc.user_id = $1", productsInCartTable, productsTable)
	if err := r.db.Select(&productsInCart, query, userId); err != nil {
		return nil, err
	}
	return productsInCart, nil
}

func (r *CartPostgres) DeleteProductFromCart(productId int, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE product_id = $1 AND user_id = $2", productsInCartTable)
	_, err := r.db.Exec(query, productId, userId)
	return err
}

func (r *CartPostgres) CleanUserCart(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", productsInCartTable)
	_, err := r.db.Exec(query, userId)
	return err
}
