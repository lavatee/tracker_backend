package model

type Order struct {
	Id              int              `json:"id" db:"id"`
	UserId          int              `json:"user_id" db:"user_id"`
	Status          string           `json:"status" db:"status"`
	OrderPrice      int              `json:"order_price" db:"order_price"`
	OrderedProducts []OrderedProduct `json:"ordered_products" db:"ordered_products"`
}

type OrderedProduct struct {
	Id              int    `json:"id" db:"id"`
	OrderId         int    `json:"order_id" db:"order_id"`
	ProductId       int    `json:"product_id" db:"product_id"`
	ProductAmount   int    `json:"product_amount" db:"product_amount"`
	Price           int    `json:"price" db:"price"`
	ProductName     string `json:"product_name" db:"product_name"`
	ProductPhotoUrl string `json:"product_photo_url" db:"product_photo_url"`
	UserFirstName   string `json:"user_first_name" db:"user_first_name"`
	UserLastName    string `json:"user_last_name" db:"user_last_name"`
	UserGrade       int    `json:"user_grade" db:"user_grade"`
	UserClassLetter string `json:"user_class_letter" db:"user_class_letter"`
}
