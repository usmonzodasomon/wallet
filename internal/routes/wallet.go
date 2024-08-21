package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github/usmonzodasomon/wallet/internal/controllers"
	"github/usmonzodasomon/wallet/internal/repository"
	"github/usmonzodasomon/wallet/internal/service"
)

func walletRoutes(r *chi.Mux, db *sqlx.DB) {
	walletRepo := repository.NewWalletRepo(db)
	walletService := service.NewWalletService(walletRepo)
	_ = controllers.NewWalletController(walletService)

}
