package main

import (
	"log"

	"github.com/car-journal/config"
	"github.com/car-journal/transport/container"
	"github.com/car-journal/transport/http/server"
	"github.com/joho/godotenv"
)

func main() {
	envDir := config.Get(config.ENV_DIR)
	err := godotenv.Load(envDir)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	} else {
		http := server.CreateHttpServer(container.CreateAppContainer())
		http.Serve()
	}
}
