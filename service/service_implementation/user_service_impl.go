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

type userServiceImpl struct {
	Ur repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) service.UserService {
	return &userServiceImpl{
		Ur: userRepo,
	}
}

func (u *userServiceImpl) Add(ctx context.Context, userPayload *model.NewUserRequest) (*model.GetUserResponse, errs.Error) {
	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: userPayload.Username,
		Password: userPayload.Password,
	}

	err = user.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.Ur.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return &model.GetUserResponse{
		StatusCode: http.StatusCreated,
		Message:    "create new user successfully",
		Data:       nil,
	}, nil
}
