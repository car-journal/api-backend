package main

import (
	"github.com/car-journal/api-backend/transport/container"
	"github.com/car-journal/api-backend/transport/http/server"
)

func main() {
	http := server.CreateHTTPServer(container.CreateAppContainer())
	http.Serve()
}
