package configuration

import (
	"online-store-golang/helper"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New() Config {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)
	return &configImpl{}
}
