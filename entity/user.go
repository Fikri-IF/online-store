package entity

import (
	"online-store-golang/errs"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":       u.Id,
		"username": u.Username,
		"exp":      time.Now().UTC().Add(60 * time.Second).Unix(),
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()
	return u.signToken(claims)
}
