package controllers

import (
	"net/http"
)

type WalletServiceI interface {
	Exists(userID string) (bool, error)
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
