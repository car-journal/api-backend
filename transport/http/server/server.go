package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/car-journal/config"
	"github.com/car-journal/transport"
	"github.com/car-journal/transport/container"
	"github.com/car-journal/transport/http/route"
	"github.com/rs/cors"
)

type httpServer struct {
	app container.AppContainer
}

func (hs *httpServer) Serve() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		MaxAge:           86400,
		AllowCredentials: true,
	})
	s := &http.Server{
		Addr:         config.Get(config.HTTP_PORT),
		ReadTimeout:  config.GetIfDuration(config.I_READ_TIMEOUT),
		WriteTimeout: config.GetIfDuration(config.I_WRITE_TIMEOUT),
		IdleTimeout:  config.GetIfDuration(config.I_IDLE_TIMEOUT),
		Handler:      c.Handler(route.Route(hs.app)),
	}

	serverErrCh := make(chan error)
	go func() {
		defer close(serverErrCh)
		log.Printf("Server running at port %s", config.Get(config.HTTP_PORT))
		serverErrCh <- s.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	signal.Notify(signalChan, signals...)

	select {
	case err := <-serverErrCh:
		log.Println("Server returning error: ", err)
	case sig := <-signalChan:
		signal.Reset(signals...)
		waitFor := config.GetIfDuration(config.I_WAIT_SHUTDOWN)

		log.Printf("Got '%s' signal, Stopping (Waiting for graceful shutdown: %s)\n", sig.String(), waitFor.String())

		ctx, cancel := context.WithTimeout(context.Background(), waitFor)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Println("Shutting down server returning error", err)
		} else {
			log.Println("Shutting down server")
		}
	}
}

func CreateHttpServer(app container.AppContainer) transport.ServerInterface {
	return &httpServer{
		app: app,
	}
}
