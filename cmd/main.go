package main

import (
	"github.com/car-journal/transport/container"
	"github.com/car-journal/transport/http/server"
)

func main() {
	http := server.CreateHttpServer(container.CreateAppContainer())
	http.Serve()
}
