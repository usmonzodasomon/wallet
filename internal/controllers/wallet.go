package controllers

import (
	"encoding/json"
	"errors"
	"github/usmonzodasomon/wallet/internal/models"
	"net/http"
)

type WalletServiceI interface {
	Exists(userID string) (bool, error)
	GetBalance(userID string) (float64, error)
	AddBalance(userID string, amount int64) error
	TotalDeposits(userID string) (int64, float64, error)
}

type WalletController struct {
	service WalletServiceI
}

func NewWalletController(service WalletServiceI) *WalletController {
	return &WalletController{service: service}
}

func (c *WalletController) Exists(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(XUserId).(string)
	_, err := c.service.Exists(userID)
	if err != nil {
		if errors.Is(err, models.ErrWalletNotFound) {
			Error(w, r, http.StatusNotFound, "wallet not found")
			return
		}
		Error(w, r, http.StatusInternalServerError, "internal server error")
		return
	}
	Success(w, r, http.StatusOK, map[string]string{"message": "wallet exists"})
}

func (c *WalletController) GetBalance(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(XUserId).(string)
	balance, err := c.service.GetBalance(userID)
	if err != nil {
		if errors.Is(err, models.ErrWalletNotFound) {
			Error(w, r, http.StatusNotFound, "wallet not found")
			return
		}
		Error(w, r, http.StatusInternalServerError, "internal server error")
		return
	}
	Success(w, r, http.StatusOK, map[string]float64{"balance": balance})
}

func (c *WalletController) AddBalance(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(XUserId).(string)

	var req models.AddBalanceReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		Error(w, r, http.StatusBadRequest, "invalid request body")
		return
	}
	if err := req.Validate(); err != nil {
		Error(w, r, http.StatusBadRequest, "invalid amount")
		return
	}

	amount, err := req.AmountInt()
	if err != nil {
		Error(w, r, http.StatusBadRequest, "invalid amount")
		return
	}

	err = c.service.AddBalance(userID, amount)
	if err != nil {
		if errors.Is(err, models.ErrWalletNotFound) {
			Error(w, r, http.StatusNotFound, "wallet not found")
			return
		}
		Error(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	Success(w, r, http.StatusOK, map[string]string{"message": "balance added"})
}

func (c *WalletController) TotalDeposits(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(XUserId).(string)

	totalCount, totalSum, err := c.service.TotalDeposits(userID)
	if err != nil {
		if errors.Is(err, models.ErrWalletNotFound) {
			Error(w, r, http.StatusNotFound, "wallet not found")
			return
		}
		Error(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	Success(w, r, http.StatusOK, map[string]interface{}{"total_count": totalCount, "total_sum": totalSum})
}
