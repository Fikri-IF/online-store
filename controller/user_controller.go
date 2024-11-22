package controller

import (
	"online-store-golang/errs"
	"online-store-golang/model"
	"online-store-golang/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
}

type userControllerImpl struct {
	Us service.UserService
}

func NewUserController(userController service.UserService) UserController {
	return &userControllerImpl{
		Us: userController,
	}
}

func (u *userControllerImpl) Register(ctx *gin.Context) {
	userPayload := &model.NewUserRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson.Error())
		return
	}

	response, err := u.Us.Add(ctx, userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
