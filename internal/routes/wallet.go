package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github/usmonzodasomon/wallet/internal/controllers"
	"github/usmonzodasomon/wallet/internal/repository"
	"github/usmonzodasomon/wallet/internal/service"
	"log/slog"
)

func walletRoutes(r *chi.Mux, db *sqlx.DB, logger *slog.Logger) {
	walletRepo := repository.NewWalletRepo(db)
	walletService := service.NewWalletService(walletRepo)
	walletControllers := controllers.NewWalletController(logger, walletService)

	r.Route("/api/v1/wallets", func(r chi.Router) {
		r.Post("/exists", walletControllers.Exists)
		r.Post("/deposit", walletControllers.AddBalance)
		r.Post("/total-deposits", walletControllers.TotalDeposits)
		r.Post("/balance", walletControllers.GetBalance)
	})
}
