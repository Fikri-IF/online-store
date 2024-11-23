package model

type UserRequest struct {
	Username string `json:"username" valid:"required~Username can't be empty" example:"monday"`
	Password string `json:"password" valid:"required~Password can't be empty, length(6|255)~Minimum password characters are 6 characters" example:"secret"`
}
