package service

import (
	"context"
	"online-store-golang/errs"
	"online-store-golang/model"
)

type CategoryService interface {
	Create(ctx context.Context, category *model.CreateCategoryRequest) (*model.GeneralResponse, errs.Error)
	FindAll(ctx context.Context) (*model.GetAllCategoryResponse, errs.Error)
	Delete(ctx context.Context, id int) (*model.GeneralResponse, errs.Error)
}
