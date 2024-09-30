package main

import (
	"log"

	gorestapi "github.com/ingerstep/go-restapi"
)

func main() {
	srv := new(gorestapi.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
