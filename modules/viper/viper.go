package viper

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func NewViper() *viper.Viper {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Load .env from root directory
	if err := godotenv.Load(filepath.Join(currentDir, ".env")); err != nil {
		panic(err)
	}

	v := viper.New()
	v.AutomaticEnv()

	env := os.Getenv("APP_ENV")
	if env == "" {
		panic(err)
	}

	v.SetConfigName(fmt.Sprintf("config-%s", env))
	v.SetConfigType("yaml")
	v.AddConfigPath(filepath.Join(currentDir, "/modules/config"))

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	return v
}
