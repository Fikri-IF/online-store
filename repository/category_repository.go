package repository

import (
	"context"
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/model"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *entity.Category) errs.Error
	FindById(ctx context.Context, id int) (*entity.Category, errs.Error)
	FindAll(ctx context.Context) ([]model.CategoryResponse, errs.Error)
}
