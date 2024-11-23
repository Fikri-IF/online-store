package model

type GeneralResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
