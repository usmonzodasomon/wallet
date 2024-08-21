package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func SetUpRoutes(r *chi.Mux, db *sqlx.DB) {
	walletRoutes(r, db)
}
