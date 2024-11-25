package service

import (
	"context"
	"online-store-golang/errs"
	"online-store-golang/model"
)

type ProductService interface {
	Create(ctx context.Context, product *model.ProductRequest) (*model.GeneralResponse, errs.Error)
	FindAll(ctx context.Context) (*model.GetAllProductsResponse, errs.Error)
	FindByCategory(ctx context.Context, id int) (*model.GetAllProductsResponse, errs.Error)
}
