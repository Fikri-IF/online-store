package service

import (
	"context"
	"online-store-golang/errs"
	"online-store-golang/model"
)

type UserService interface {
	Add(ctx context.Context, userPayload *model.NewUserRequest) (*model.GetUserResponse, errs.Error)
}
