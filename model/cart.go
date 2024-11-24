package model

type AddToCartRequest struct {
	ProductId int `json:"product_id" valid:"required"`
	Quantity  int `json:"quantity"`
}

type CartItemResponse struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
