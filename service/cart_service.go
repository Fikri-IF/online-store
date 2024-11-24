package service

import (
	"context"
	"online-store-golang/errs"
	"online-store-golang/model"
)

type CartService interface {
	AddItem(ctx context.Context, userId int, addToCartPayload *model.AddToCartRequest) (*model.GeneralResponse, errs.Error)
	GetUserCart(ctx context.Context, userId int) (*model.GeneralResponse, errs.Error)
}
