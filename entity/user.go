package entity

import (
	"online-store-golang/errs"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HashPassword() errs.Error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	u.Password = string(hashPassword)

	return nil
}
