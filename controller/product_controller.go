package controller

import (
	"online-store-golang/errs"
	"online-store-golang/helper"
	"online-store-golang/model"
	"online-store-golang/service"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAll(ctx *gin.Context)
	FindByCategory(ctx *gin.Context)
	Create(ctx *gin.Context)
	FindById(ctx *gin.Context)
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
	categoryId, errParam := helper.GetParamId(ctx, "categoryId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}
	response, err := p.Ps.FindByCategory(ctx, categoryId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
func (p *ProductControllerImpl) Create(ctx *gin.Context) {
	productPayload := &model.ProductRequest{}

	if err := ctx.ShouldBindJSON(productPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson.Error())
		return
	}

	response, err := p.Ps.Create(ctx, productPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
func (p *ProductControllerImpl) FindById(ctx *gin.Context) {
	productId, errParam := helper.GetParamId(ctx, "productId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}
	response, err := p.Ps.FindById(ctx, productId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
