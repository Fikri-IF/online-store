package serviceimplementation

import (
	"context"
	"net/http"
	"online-store-golang/errs"
	"online-store-golang/model"
	"online-store-golang/repository"
	"online-store-golang/service"
)

type productServiceImpl struct {
	Pr repository.ProductRepository
}

func NewProductService(ProductRepository repository.ProductRepository) service.ProductService {
	return &productServiceImpl{
		Pr: ProductRepository,
	}
}

func (p *productServiceImpl) FindAll(ctx context.Context) (*model.GetAllProductsResponse, errs.Error) {
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

func (p *productServiceImpl) FindByCategory(ctx context.Context, id int) (*model.GetAllProductsResponse, errs.Error) {
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
