package authservice

import (
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type AuthServiceImpl struct {
	Ur repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		Ur: userRepo,
	}
}

func (a *AuthServiceImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		invalidToken := errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		user := entity.User{}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.Ur.FetchById(ctx, user.Id)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidToken.Status(), invalidToken)
			return
		}

		ctx.Set("userData", user)
		ctx.Next()
	}
}
