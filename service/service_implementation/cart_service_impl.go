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

type CartServiceImpl struct {
	Cr repository.CartRepository
}

func NewCartService(CartRepository repository.CartRepository) service.CartService {
	return &CartServiceImpl{
		Cr: CartRepository,
	}
}

func (cs *CartServiceImpl) AddItem(ctx context.Context, userId int, addToCartPayload *model.AddToCartRequest) (*model.GeneralResponse, errs.Error) {
	err := helper.ValidateStruct(addToCartPayload)
	if err != nil {
		return nil, err
	}

	cartItem := &entity.CartItem{
		UserId:    userId,
		ProductId: addToCartPayload.ProductId,
		Quantity:  addToCartPayload.Quantity,
	}

	err = cs.Cr.AddItem(ctx, cartItem)

	if err != nil {
		return nil, err
	}

	return &model.GeneralResponse{
		StatusCode: http.StatusCreated,
		Message:    "new item successfully added",
		Data:       nil,
	}, nil
}
