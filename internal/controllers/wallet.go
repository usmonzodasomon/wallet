package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/middleware"
	"github/usmonzodasomon/wallet/internal/models"
	"log/slog"
	"net/http"
)

type WalletServiceI interface {
	Exists(userID string) (bool, error)
	GetBalance(userID string) (float64, error)
	AddBalance(userID string, amount int64) error
	TotalDeposits(userID string) (int64, float64, error)
}

type WalletController struct {
	logger  *slog.Logger
	service WalletServiceI
}

func NewWalletController(logger *slog.Logger, service WalletServiceI) *WalletController {
	return &WalletController{logger: logger, service: service}
}

func (c *WalletController) Exists(w http.ResponseWriter, r *http.Request) {
	log := c.logger.With(
		slog.String("fn", "Exists"),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	log.Info("checking if wallet exists")

	userID := r.Context().Value(XUserId).(string)
	_, err := c.service.Exists(userID)
	if err != nil {
		if errors.Is(err, models.ErrWalletNotFound) {
			log.Warn("wallet not found")
			Error(w, r, http.StatusNotFound, "wallet not found")
			return
		}
		log.Error("internal server error", slog.String("error", err.Error()))
		Error(w, r, http.StatusInternalServerError, "internal server error")
		return
	}
	log.Info("wallet exists")
	Success(w, r, http.StatusOK, map[string]string{"message": "wallet exists"})
}

func (c *WalletController) GetBalance(w http.ResponseWriter, r *http.Request) {
	log := c.logger.With(
		slog.String("fn", "GetBalance"),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	log.Info("getting balance")

	userID := r.Context().Value(XUserId).(string)
	balance, err := c.service.GetBalance(userID)
	if err != nil {
		if errors.Is(err, models.ErrWalletNotFound) {
			log.Warn("wallet not found")
			Error(w, r, http.StatusNotFound, "wallet not found")
			return
		}
		log.Error("internal server error", slog.String("error", err.Error()))
		Error(w, r, http.StatusInternalServerError, "internal server error")
		return
	}
	log.Info("balance received")
	Success(w, r, http.StatusOK, map[string]float64{"balance": balance})
}

func (c *WalletController) AddBalance(w http.ResponseWriter, r *http.Request) {
	log := c.logger.With(
		slog.String("fn", "Exists"),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	log.Info("adding balance")
	userID := r.Context().Value(XUserId).(string)

	var req models.AddBalanceReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error("invalid request body", slog.String("error", err.Error()))
		Error(w, r, http.StatusBadRequest, "invalid request body")
		return
	}
	if err := req.Validate(); err != nil {
		log.Warn("invalid amount")
		Error(w, r, http.StatusBadRequest, "invalid amount")
		return
	}

	amount, err := req.AmountInt()
	if err != nil {
		log.Warn("invalid amount")
		Error(w, r, http.StatusBadRequest, "invalid amount")
		return
	}

	err = c.service.AddBalance(userID, amount)
	if err != nil {
		if errors.Is(err, models.ErrWalletNotFound) {
			log.Warn("wallet not found")
			Error(w, r, http.StatusNotFound, "wallet not found")
			return
		} else if errors.Is(err, models.ErrMaxBalanceExceeded) {
			log.Warn("max balance exceeded")
			Error(w, r, http.StatusBadRequest, "max balance exceeded")
			return
		}
		log.Error("internal server error", slog.String("error", err.Error()))
		Error(w, r, http.StatusInternalServerError, "internal server error")
		return
	}
	log.Info("balance added")
	Success(w, r, http.StatusOK, map[string]string{"message": "balance added"})
}

func (c *WalletController) TotalDeposits(w http.ResponseWriter, r *http.Request) {
	log := c.logger.With(
		slog.String("fn", "Exists"),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	log.Info("getting total deposits")
	userID := r.Context().Value(XUserId).(string)

	totalCount, totalSum, err := c.service.TotalDeposits(userID)
	if err != nil {
		if errors.Is(err, models.ErrWalletNotFound) {
			log.Warn("wallet not found")
			Error(w, r, http.StatusNotFound, "wallet not found")
			return
		}
		log.Error("internal server error", slog.String("error", err.Error()))
		Error(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	log.Info("total deposits received")
	Success(w, r, http.StatusOK, map[string]interface{}{"total_count": totalCount, "total_sum": totalSum})
}
