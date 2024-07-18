package cfg

import (
	"errors"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var ErrInvalidAppName = errors.New("app name is invalid")

// LoadWithAPS loads configuration from local .env file and AWS Parameter Store as well
func LoadWithAPS(out interface{}, stage string) error {
	switch stage {
	case "", "local":
		if err := PreloadLocalENV(); err != nil {
			return err
		}
	case "dev":
		return nil
	}

	return env.Parse(out)
}

// PreloadLocalENV reads .env* files and sets the values to os ENV
func PreloadLocalENV() error {
	return godotenv.Load(".env")
}
