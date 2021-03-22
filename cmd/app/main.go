package main

import (
	"log"

	server "github.com/Valeriy-Totubalin/test_project"
	"github.com/Valeriy-Totubalin/test_project/internal/app/config"
	"github.com/Valeriy-Totubalin/test_project/internal/delivery/handler"
)

func main() {
	config.Init()
	serverConfig := config.NewServer()
	mysqlConfig := config.NewDBMysql()
	conf := config.NewConfig(
		mysqlConfig,
		serverConfig,
	)

	srv := &server.Server{
		Config: conf.Srv(),
	}

	handlers := new(handler.Handler)

	if err := srv.Run(handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}
