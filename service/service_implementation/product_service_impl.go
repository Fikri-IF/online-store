package serviceimplementation

import (
	"context"
	"net/http"
	"online-store-golang/errs"
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
