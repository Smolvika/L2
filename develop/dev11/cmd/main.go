package main

import (
	"context"
	"develop/dev11"
	"develop/dev11/pkg/handler"
	"develop/dev11/pkg/repository"
	"develop/dev11/pkg/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandlers(services)
	srv := new(dev11.Server)
	if err := srv.Run("8080", handlers.InitRouters()); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error running http server: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}
