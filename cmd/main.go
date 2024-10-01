package main

import (
	"log"

	gorestapi "github.com/ingerstep/go-restapi"
	"github.com/ingerstep/go-restapi/pkg/handler"
	"github.com/ingerstep/go-restapi/pkg/repository"
	"github.com/ingerstep/go-restapi/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		Username: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("POSTGRES_DB"),
		SSLMode:  viper.GetString("SSL_MODE"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gorestapi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("env")
	viper.SetConfigFile("../.env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
