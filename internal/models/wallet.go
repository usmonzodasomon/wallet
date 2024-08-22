package models

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrWalletNotFound = errors.New("wallet not found")
	ErrInvalidAmount  = errors.New("invalid amount")
)

type Wallet struct {
	ID      uint64
	Balance int64
	UserID  string
}

type AddBalanceReq struct {
	Amount string `json:"amount"`
}

func (r AddBalanceReq) Validate() error {
	if r.Amount == "" {
		return ErrInvalidAmount
	}

	parts := strings.Split(r.Amount, ".")
	if len(parts) != 2 {
		return ErrInvalidAmount
	}

	return nil
}

func (r AddBalanceReq) AmountInt() (int64, error) {
	if err := r.Validate(); err != nil {
		return 0, err
	}
	parts := strings.Split(r.Amount, ".")
	somoni, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, ErrInvalidAmount
	}
	diram, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, ErrInvalidAmount
	}
	return somoni*100 + diram, nil
}
