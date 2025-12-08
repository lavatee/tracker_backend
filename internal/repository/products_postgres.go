package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lavatee/tracker_backend/internal/model"
)

type ProductsPostgres struct {
	db *sqlx.DB
}

func NewProductsPostgres(db *sqlx.DB) *ProductsPostgres {
	return &ProductsPostgres{
		db: db,
	}
}

func (r *ProductsPostgres) CreateProduct(product model.Product) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description, price, photo_url) VALUES ($1, $2, $3, $4) RETURNING id", productsTable)
	row := r.db.QueryRow(query, product.Name, product.Description, product.Price, product.PhotoUrl)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductsPostgres) GetProducts() ([]model.Product, error) {
	var products []model.Product
	query := fmt.Sprintf("SELECT * FROM %s", productsTable)
	if err := r.db.Select(&products, query); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductsPostgres) GetProductById(productId int) (model.Product, error) {
	var product model.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", productsTable)
	if err := r.db.Get(&product, query, productId); err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *ProductsPostgres) DeleteProduct(productId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", productsTable)
	_, err := r.db.Exec(query, productId)
	return err
}

func (r *ProductsPostgres) UpdateProduct(product model.Product) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, description = $2, price = $3, photo_url = $4 WHERE id = $5", productsTable)
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.PhotoUrl, product.Id)
	return err
}
