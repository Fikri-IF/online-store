package controller

import (
	"online-store-golang/helper"
	"online-store-golang/service"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAll(ctx *gin.Context)
	FindByCategory(ctx *gin.Context)
}

type ProductControllerImpl struct {
	Ps service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		Ps: productService,
	}
}

func (p *ProductControllerImpl) FindAll(ctx *gin.Context) {
	response, err := p.Ps.FindAll(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}

func (p *ProductControllerImpl) FindByCategory(ctx *gin.Context) {
	photoId, errParam := helper.GetParamId(ctx, "categoryId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}
	response, err := p.Ps.FindByCategory(ctx, photoId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
