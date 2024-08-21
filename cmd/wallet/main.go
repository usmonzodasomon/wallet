package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github/usmonzodasomon/wallet/internal/config"
	"github/usmonzodasomon/wallet/internal/routes"
	"github/usmonzodasomon/wallet/pkg/logger"
	"github/usmonzodasomon/wallet/pkg/postgres"
	"github/usmonzodasomon/wallet/pkg/server"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.Logger(cfg.Env)

	logger.Info("starting url-shortener",
		slog.String("env", cfg.Env))

	logger.Debug("debug messages are enabled")

	connection, err := postgres.GetConnection(postgres.Config{
		Host:     cfg.Database.PostgresHost,
		Port:     cfg.Database.PostgresPort,
		User:     cfg.Database.PostgresUser,
		Password: cfg.Database.PostgresPassword,
		Database: cfg.Database.PostgresDatabase,
	})
	if err != nil {
		logger.Error("Failed to connect to database: ", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer postgres.CloseConnection(connection)
	logger.Info("connected to database")

	r := chi.NewRouter()
	routes.SetUpRoutes(r, connection)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	logger.Info("starting server", slog.String("address", cfg.Address))

	srv := server.Server{}
	go func() {
		if err := srv.Run(server.Config{
			Address:      cfg.HTTPServer.Address,
			ReadTimeout:  cfg.HTTPServer.Timeout,
			WriteTimeout: cfg.HTTPServer.Timeout,
			IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		}, r); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("failed to start server")
		}
	}()

	logger.Info(fmt.Sprintf("server started on %s", cfg.Address))
	<-done
	logger.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("failed to stop server", slog.String("error", err.Error()))
		return
	}
	logger.Info("server stopped")
}
