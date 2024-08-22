package controllers

import (
	"errors"
	"github/usmonzodasomon/wallet/internal/models"
	"net/http"
)

type WalletServiceI interface {
	Exists(userID string) (bool, error)
	GetBalance(userID string) (float64, error)
}

type WalletController struct {
	service WalletServiceI
}

func NewWalletController(service WalletServiceI) *WalletController {
	return &WalletController{service: service}
}

func (c *WalletController) Exists(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(XUserId).(string)
	exists, err := c.service.Exists(userID)
	if err != nil {
		Error(w, r, http.StatusInternalServerError, "internal server error")
		return
	}
	if !exists {
		Error(w, r, http.StatusNotFound, "wallet not found")
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
