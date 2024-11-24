package repository

import (
	"context"
	"online-store-golang/entity"
	"online-store-golang/errs"
)

type UserRepository interface {
	Create(ctx context.Context, userPayload *entity.User) errs.Error
	FetchByUsername(ctx context.Context, username string) (*entity.User, errs.Error)
	FetchById(ctx context.Context, id int) (*entity.User, errs.Error)
}
