package controllers

import (
	"encoding/json"
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
	userID := r.Context().Value("X-UserId").(string)
	exists, err := c.service.Exists(userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "wallet not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "wallet exists"})
}
