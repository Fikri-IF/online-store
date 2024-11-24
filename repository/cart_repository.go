package repository

import (
	"context"
	"online-store-golang/entity"
	"online-store-golang/errs"
)

type CartRepository interface {
	AddItem(ctx context.Context, cartItemPayload *entity.CartItem) errs.Error
}
