package main

import (
	"github.com/raufhm/golang-project-convention/modules/config"
	"github.com/raufhm/golang-project-convention/modules/logger"
	"github.com/raufhm/golang-project-convention/modules/server"
	"github.com/raufhm/golang-project-convention/modules/viper"
	"github.com/raufhm/golang-project-convention/services/iam/features/health"
)

func main() {
	log := logger.NewLogger()
	vpr := viper.NewViper()
	cfg := config.NewConfig(vpr)
	srv := server.NewServer(cfg)

	healthService := health.NewService(log)
	healthHandler := health.NewHandler(healthService)
	healthHandler.Register(srv.GetEcho())

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}

}
