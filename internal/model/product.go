package model

type Product struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Price       int    `json:"price" db:"price"`
	PhotoUrl    string `json:"photo_url" db:"photo_url"`
	Description string `json:"description" db:"description"`
}

type ProductInCart struct {
	Id              int    `json:"id" db:"id"`
	UserId          int    `json:"user_id" db:"user_id"`
	ProductId       int    `json:"product_id" db:"product_id"`
	ProductAmount   int    `json:"product_amount" db:"product_amount"`
	ProductName     string `json:"product_name" db:"product_name"`
	ProductPhotoUrl string `json:"product_photo_url" db:"product_photo_url"`
}
