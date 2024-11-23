package service

import (
	"context"
	"online-store-golang/errs"
	"online-store-golang/model"
)

type UserService interface {
	Add(ctx context.Context, userPayload *model.UserRequest) (*model.GeneralResponse, errs.Error)
	Login(ctx context.Context, userPayload *model.UserRequest) (*model.GeneralResponse, errs.Error)
}
