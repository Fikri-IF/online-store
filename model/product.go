package model

type GetProductResponse struct {
	ProductId   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Stock       int    `json:"stock"`
}

type GetAllProductsResponse struct {
	StatusCode int                  `json:"status_code"`
	Message    string               `json:"message"`
	Products   []GetProductResponse `json:"products"`
}

type ProductRequest struct {
	ProductName string `json:"product_name" valid:"required~Product name is blank"`
	Price       int    `json:"price" valid:"int, required~Price is not valid"`
	Category    int    `json:"category" valid:"int~Category is not valid"`
	Stock       int    `json:"stock" valid:"required~Stock is not valid"`
}
