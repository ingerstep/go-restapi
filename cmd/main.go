package main

import (
	"log"

	gorestapi "github.com/ingerstep/go-restapi"
	"github.com/ingerstep/go-restapi/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(gorestapi.Server)
	if err := srv.Run("3005", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
