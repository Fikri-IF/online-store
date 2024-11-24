package model

type AddToCartRequest struct {
	ProductId int `json:"product_id" valid:"required"`
	Quantity  int `json:"quantity"`
}

type CartItemResponse struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type UserCartResponse struct {
	ProductId   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}
