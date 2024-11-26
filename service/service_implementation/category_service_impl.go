package serviceimplementation

import (
	"context"
	"net/http"
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/helper"
	"online-store-golang/model"
	"online-store-golang/repository"
	"online-store-golang/service"
)

type CategoryServiceImpl struct {
	Cr repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) service.CategoryService {
	return &CategoryServiceImpl{
		Cr: categoryRepository,
	}
}

func (cs *CategoryServiceImpl) Create(ctx context.Context, category *model.CreateCategoryRequest) (*model.GeneralResponse, errs.Error) {
	err := helper.ValidateStruct(category)
	if err != nil {
		return nil, err
	}
	categoryEntity := &entity.Category{
		Name: category.Name,
	}
	err = cs.Cr.Create(ctx, categoryEntity)
	if err != nil {
		return nil, err
	}
	return &model.GeneralResponse{
		StatusCode: http.StatusCreated,
		Message:    "category successfully created",
		Data:       categoryEntity,
	}, nil
}
func (cs *CategoryServiceImpl) FindAll(ctx context.Context) (*model.GetAllCategoryResponse, errs.Error) {
	categories, err := cs.Cr.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return &model.GetAllCategoryResponse{
		StatusCode: http.StatusOK,
		Message:    "categories successfully fetched",
		Categories: categories,
	}, nil
}

func (cs *CategoryServiceImpl) Delete(ctx context.Context, id int) (*model.GeneralResponse, errs.Error) {
	category, err := cs.Cr.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	err = cs.Cr.Delete(ctx, category.CategoryId)
	if err != nil {
		return nil, err
	}
	return &model.GeneralResponse{
		StatusCode: http.StatusOK,
		Message:    "category successfully deleted",
		Data:       nil,
	}, nil
}
