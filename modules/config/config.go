package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database Database
	Server   Server
}

type Server struct {
	Port int64
	Host string
}

type Database struct {
	DatabaseURL string
}

func NewConfig(v *viper.Viper) *Config {
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}
	config.Server.Port = v.GetInt64("app.port")

	return &config
}
