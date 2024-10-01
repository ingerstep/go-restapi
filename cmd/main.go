package main

import (
	"log"

	gorestapi "github.com/ingerstep/go-restapi"
	"github.com/ingerstep/go-restapi/pkg/handler"
	"github.com/ingerstep/go-restapi/pkg/repository"
	"github.com/ingerstep/go-restapi/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initCongif(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gorestapi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initCongif() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
