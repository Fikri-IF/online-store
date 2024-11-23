package helper

import (
	"online-store-golang/errs"

	"github.com/asaskevich/govalidator"
)

func ValidateStruct(s interface{}) errs.Error {

	_, err := govalidator.ValidateStruct(s)

	if err != nil {
		return errs.NewBadRequestError(err.Error())
	}

	return nil
}