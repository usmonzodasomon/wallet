package routes

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github/usmonzodasomon/wallet/internal/controllers/middlewares"
	"log/slog"
)

func SetUpRoutes(r *chi.Mux, db *sqlx.DB, logger *slog.Logger) {
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middlewares.CheckHashMiddleware)

	walletRoutes(r, db, logger)
}
