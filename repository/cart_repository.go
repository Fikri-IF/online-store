package repository

import (
	"context"
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/model"
)

type CartRepository interface {
	AddItem(ctx context.Context, cartItemPayload *entity.CartItem) errs.Error
	GetUserCart(ctx context.Context, userId int) ([]model.UserCartResponse, errs.Error)
	GetItem(ctx context.Context, userId int, productId int) (*entity.CartItem, errs.Error)
	UpdateQuantity(ctx context.Context, cartItemPayload *entity.CartItem) errs.Error
}
