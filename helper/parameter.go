package helper

import (
	"fmt"
	"online-store-golang/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamId(c *gin.Context, key string) (int, errs.Error) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, errs.NewBadRequestError(fmt.Sprintf("parameter '%s' has to be a number", key))
	}

	return id, nil
}
