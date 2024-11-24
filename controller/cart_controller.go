package controller

import (
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/helper"
	"online-store-golang/model"
	"online-store-golang/service"

	"github.com/gin-gonic/gin"
)

type CartController interface {
	AddToCart(ctx *gin.Context)
	GetUserCart(ctx *gin.Context)
	DeleteItem(ctx *gin.Context)
}

type CartControllerImpl struct {
	Cs service.CartService
}

func NewCartController(cartService service.CartService) CartController {
	return &CartControllerImpl{
		Cs: cartService,
	}
}

func (c *CartControllerImpl) AddToCart(ctx *gin.Context) {
	user, ok := ctx.MustGet("userData").(entity.User)

	if !ok {
		internalServerErr := errs.NewInternalServerError("something went wrong")
		ctx.AbortWithStatusJSON(internalServerErr.Status(), internalServerErr)
	}

	cartItemPayload := &model.AddToCartRequest{}

	if err := ctx.ShouldBindJSON(cartItemPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson.Error())
		return
	}

	response, err := c.Cs.AddItem(ctx, user.Id, cartItemPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
func (c *CartControllerImpl) GetUserCart(ctx *gin.Context) {
	user, ok := ctx.MustGet("userData").(entity.User)

	if !ok {
		internalServerErr := errs.NewInternalServerError("something went wrong")
		ctx.AbortWithStatusJSON(internalServerErr.Status(), internalServerErr)
	}

	response, err := c.Cs.GetUserCart(ctx, user.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}

func (c *CartControllerImpl) DeleteItem(ctx *gin.Context) {
	user, ok := ctx.MustGet("userData").(entity.User)

	productId, errParam := helper.GetParamId(ctx, "productId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	if !ok {
		internalServerErr := errs.NewInternalServerError("something went wrong")
		ctx.AbortWithStatusJSON(internalServerErr.Status(), internalServerErr)
	}
	response, err := c.Cs.DeleteItem(ctx, user.Id, productId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
