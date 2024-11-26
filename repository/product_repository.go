package repository

import (
	"context"
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/model"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) errs.Error
	FindAll(ctx context.Context) ([]model.GetProductResponse, errs.Error)
	FindByCategory(ctx context.Context, id int) ([]model.GetProductResponse, errs.Error)
	FindById(ctx context.Context, id int) (*model.GetProductResponse, errs.Error)
	Delete(ctx context.Context, id int) errs.Error
}
