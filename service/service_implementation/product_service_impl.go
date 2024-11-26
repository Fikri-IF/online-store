package serviceimplementation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/helper"
	"online-store-golang/model"
	"online-store-golang/repository"
	"online-store-golang/service"
	"time"

	"github.com/redis/go-redis/v9"
)

type ProductServiceImpl struct {
	Pr    repository.ProductRepository
	Cr    repository.CategoryRepository
	Cache *redis.Client
}

func NewProductService(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository, cache *redis.Client) service.ProductService {
	return &ProductServiceImpl{
		Pr:    productRepository,
		Cr:    categoryRepository,
		Cache: cache,
	}
}

func (p *ProductServiceImpl) Create(ctx context.Context, product *model.ProductRequest) (*model.GeneralResponse, errs.Error) {
	err := helper.ValidateStruct(product)
	if err != nil {
		return nil, err
	}
	_, err = p.Cr.FindById(ctx, product.Category)

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
		Data:       productEntity,
	}, nil
}

func (p *ProductServiceImpl) FindAll(ctx context.Context) (*model.GetAllProductsResponse, errs.Error) {
	var products []model.GetProductResponse

	cacheKey := "all_products"

	cachedProducts, redisErr := p.Cache.Get(ctx, cacheKey).Result()
	if redisErr == nil {
		redisErr = json.Unmarshal([]byte(cachedProducts), &products)
		if redisErr != nil {
			return nil, errs.NewInternalServerError("failed to unmarshal cached data")
		}
		return &model.GetAllProductsResponse{
			StatusCode: http.StatusOK,
			Message:    "products successfully fetched",
			Products:   products,
		}, nil
	}
	products, err := p.Pr.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	cachedData, redisErr := json.Marshal(products)
	if redisErr != nil {
		return nil, errs.NewInternalServerError("failed to marshal products")
	}
	if redisErr := p.Cache.Set(ctx, cacheKey, cachedData, time.Hour).Err(); redisErr != nil {
		return nil, errs.NewInternalServerError("failed to set cache")
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
func (p *ProductServiceImpl) FindById(ctx context.Context, productId int) (*model.GeneralResponse, errs.Error) {
	var product *model.GetProductResponse

	cacheKey := fmt.Sprintf("product:details:%d", productId)

	cachedProduct, redisErr := p.Cache.Get(ctx, cacheKey).Result()
	if redisErr == nil {
		redisErr = json.Unmarshal([]byte(cachedProduct), &product)
		if redisErr != nil {
			return nil, errs.NewInternalServerError("failed to unmarshal cached data")
		}
		return &model.GeneralResponse{
			StatusCode: http.StatusOK,
			Message:    "product found",
			Data:       product,
		}, nil
	}

	product, err := p.Pr.FindById(ctx, productId)
	if err != nil {
		return nil, err
	}

	cachedData, redisErr := json.Marshal(product)
	if redisErr != nil {
		return nil, errs.NewInternalServerError("failed to marshal products")
	}
	if redisErr := p.Cache.Set(ctx, cacheKey, cachedData, time.Hour).Err(); redisErr != nil {
		return nil, errs.NewInternalServerError("failed to set cache")
	}

	return &model.GeneralResponse{
		StatusCode: http.StatusOK,
		Message:    "product found",
		Data:       product,
	}, nil
}
func (p *ProductServiceImpl) Delete(ctx context.Context, productId int) (*model.GeneralResponse, errs.Error) {
	product, err := p.Pr.FindById(ctx, productId)
	if err != nil {
		return nil, err
	}
	err = p.Pr.Delete(ctx, product.ProductId)
	if err != nil {
		return nil, err
	}
	return &model.GeneralResponse{
		StatusCode: http.StatusOK,
		Message:    "product successfully deleted",
		Data:       nil,
	}, nil
}
