package repository

import (
	"context"
	"online-store-golang/entity"
	"online-store-golang/errs"
)

type UserRepository interface {
	Create(ctx context.Context, userPayload *entity.User) errs.Error
}
