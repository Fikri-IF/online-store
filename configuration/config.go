package configuration

import (
	"os"
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
	// err := godotenv.Load(".env")
	// errs.PanicIfError(err)
	return &configImpl{}
}
