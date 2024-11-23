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
