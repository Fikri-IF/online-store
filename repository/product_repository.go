package repository

import (
	"context"
	"online-store-golang/errs"
	"online-store-golang/model"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]model.GetProductResponse, errs.Error)
	FindByCategory(ctx context.Context, id int) ([]model.GetProductResponse, errs.Error)
}
