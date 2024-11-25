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

type ProductServiceImpl struct {
	Pr repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) service.ProductService {
	return &ProductServiceImpl{
		Pr: productRepository,
	}
}

func (p *ProductServiceImpl) Create(ctx context.Context, product *model.ProductRequest) (*model.GeneralResponse, errs.Error) {
	err := helper.ValidateStruct(product)
	if err != nil {
		return nil, err
	}

	productEntity := &entity.Product{
		ProductName: product.ProductName,
		Price:       product.Price,
		Category:    product.Category,
		Stock:       product.Stock,
	}
	err = p.Pr.Create(ctx, productEntity)
	if err != nil {
		return nil, err
	}
	return &model.GeneralResponse{
		StatusCode: http.StatusCreated,
		Message:    "product successfully created",
		Data:       &productEntity,
	}, nil
}

func (p *ProductServiceImpl) FindAll(ctx context.Context) (*model.GetAllProductsResponse, errs.Error) {
	products, err := p.Pr.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return &model.GetAllProductsResponse{
		StatusCode: http.StatusOK,
		Message:    "products successfully fetched",
		Products:   products,
	}, nil
}

func (p *ProductServiceImpl) FindByCategory(ctx context.Context, id int) (*model.GetAllProductsResponse, errs.Error) {
	products, err := p.Pr.FindByCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.GetAllProductsResponse{
		StatusCode: http.StatusOK,
		Message:    "products successfully fetched",
		Products:   products,
	}, nil
}
