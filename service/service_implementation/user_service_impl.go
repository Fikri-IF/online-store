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

func (u *userServiceImpl) Add(ctx context.Context, userPayload *model.UserRequest) (*model.GeneralResponse, errs.Error) {
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

	return &model.GeneralResponse{
		StatusCode: http.StatusCreated,
		Message:    "create new user successfully",
		Data:       nil,
	}, nil
}

func (u *userServiceImpl) Login(ctx context.Context, userPayload *model.UserRequest) (*model.GeneralResponse, errs.Error) {
	err := helper.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	user, err := u.Ur.FetchByUsername(ctx, userPayload.Username)

	if err != nil {
		return nil, err
	}

	isValidPassword := user.ComparePassword(userPayload.Password)

	if !isValidPassword {
		return nil, errs.NewBadRequestError("invalid password")
	}

	token := user.GenerateToken()

	return &model.GeneralResponse{
		StatusCode: http.StatusOK,
		Message:    "login successfull",
		Data: &model.TokenResponse{
			Token: token,
		},
	}, nil
}
