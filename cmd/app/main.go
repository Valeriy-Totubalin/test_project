package main

import (
	"log"

	server "github.com/Valeriy-Totubalin/test_project"
	"github.com/Valeriy-Totubalin/test_project/internal/app/config"
	"github.com/Valeriy-Totubalin/test_project/internal/app/factories"
	"github.com/Valeriy-Totubalin/test_project/internal/delivery/handler"
	"github.com/Valeriy-Totubalin/test_project/pkg/link_manager"
	"github.com/Valeriy-Totubalin/test_project/pkg/token_manager"
)

// @title Todo App API
// @version 1.0
// @description API server for test application

// @host localhost:8080
// @BasePath /

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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

	serviceFactory := factories.NewServicesFactory(conf)
	tokenManager, err := token_manager.NewManager(conf.GetTokenSecret())
	if nil != err {
		log.Println(err.Error())
		return
	}
	linkManager, err := link_manager.NewManager(conf.GetLinkSecret())
	if nil != err {
		log.Println(err.Error())
		return
	}

	handlers := new(handler.Handler)
	handlers.ServiceFactory = serviceFactory
	handlers.TokenManager = tokenManager
	handlers.LinkManager = linkManager

	if err := srv.Run(handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}
