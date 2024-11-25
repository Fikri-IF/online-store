package model

type CreateCategoryRequest struct {
	Name string `valid:"required,stringlength(1|255)" json:"name"`
}

type CategoryResponse struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
type GetAllCategoryResponse struct {
	StatusCode int                `json:"status_code"`
	Message    string             `json:"message"`
	Categories []CategoryResponse `json:"categories"`
}
