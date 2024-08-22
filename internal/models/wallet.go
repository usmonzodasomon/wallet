package models

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrWalletNotFound     = errors.New("wallet not found")
	ErrInvalidAmount      = errors.New("invalid amount")
	ErrMaxBalanceExceeded = errors.New("max balance exceeded")
)

type Wallet struct {
	ID           int64  `db:"id"`
	Balance      int64  `db:"balance"`
	UserID       string `db:"user_id"`
	IsIdentified bool   `db:"is_identified"`
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

type TotalDeposits struct {
	TotalCount int64 `db:"total_count"`
	TotalSum   int64 `db:"total_sum"`
}
