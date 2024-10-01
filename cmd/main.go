package main

import (
	"log"

	gorestapi "github.com/ingerstep/go-restapi"
	"github.com/ingerstep/go-restapi/pkg/handler"
	"github.com/ingerstep/go-restapi/pkg/repository"
	"github.com/ingerstep/go-restapi/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gorestapi.Server)
	if err := srv.Run("3005", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
