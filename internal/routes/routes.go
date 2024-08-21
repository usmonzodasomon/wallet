package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github/usmonzodasomon/wallet/internal/controllers/middlewares"
)

func SetUpRoutes(r *chi.Mux, db *sqlx.DB) {
	r.Use(middlewares.CheckHashMiddleware)

	walletRoutes(r, db)
}
