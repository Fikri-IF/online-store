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

	currentItem, err := cs.Cr.GetItem(ctx, userId, addToCartPayload.ProductId)
	if err != nil {
		return nil, err
	}
	addsQuantity := addToCartPayload.Quantity
	if currentItem != nil {
		addsQuantity += currentItem.Quantity
	}

	cartItem := &entity.CartItem{
		UserId:    userId,
		ProductId: addToCartPayload.ProductId,
		Quantity:  addsQuantity,
	}
	var statusCode int

	if currentItem != nil {
		err = cs.Cr.UpdateQuantity(ctx, cartItem)
		if err != nil {
			return nil, err
		}
		statusCode = http.StatusOK
	} else {
		err = cs.Cr.AddItem(ctx, cartItem)
		if err != nil {
			return nil, err
		}
		statusCode = http.StatusCreated
	}

	return &model.GeneralResponse{
		StatusCode: statusCode,
		Message:    "item successfully added",
		Data:       nil,
	}, nil
}

func (cs *CartServiceImpl) GetUserCart(ctx context.Context, userId int) (*model.GeneralResponse, errs.Error) {

	items, err := cs.Cr.GetUserCart(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &model.GeneralResponse{
		StatusCode: http.StatusOK,
		Message:    "cart items successfully fetched",
		Data:       items,
	}, nil
}
