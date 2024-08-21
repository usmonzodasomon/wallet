package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github/usmonzodasomon/wallet/internal/routes"
	"github/usmonzodasomon/wallet/pkg/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := chi.NewRouter()
	routes.SetUpRoutes(r, nil)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	srv := server.Server{}
	go func() {
		if err := srv.Run(server.Config{}, r); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("failed to start server")
		}
	}()
	<-done
	log.Println("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err.Error())
		return
	}
}
