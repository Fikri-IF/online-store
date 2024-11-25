package controller

import (
	"online-store-golang/errs"
	"online-store-golang/model"
	"online-store-golang/service"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type CategoryControllerImpl struct {
	Cs service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Cs: categoryService,
	}
}

func (cc *CategoryControllerImpl) Create(ctx *gin.Context) {
	categoryPayload := &model.CreateCategoryRequest{}

	if err := ctx.ShouldBindJSON(categoryPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson.Error())
		return
	}

	response, err := cc.Cs.Create(ctx, categoryPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
func (cc *CategoryControllerImpl) FindAll(ctx *gin.Context) {
	response, err := cc.Cs.FindAll(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
