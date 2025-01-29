package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
)

type Params struct {
	fx.In
}

func NewViper() (*viper.Viper, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return nil, fmt.Errorf("APP_ENV is not set")
	}

	v := viper.New()
	v.SetConfigName(fmt.Sprintf("config-%s", env))
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return v, nil
}
