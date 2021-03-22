package main

import (
	"log"

	server "github.com/Valeriy-Totubalin/test_project"
	"github.com/Valeriy-Totubalin/test_project/internal/delivery/handler"
)

func main() {
	srv := new(server.Server)

	handlers := new(handler.Handler)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}
