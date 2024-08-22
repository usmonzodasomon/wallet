package models

import "errors"

var (
	ErrWalletNotFound = errors.New("wallet not found")
)

type Wallet struct {
	ID      uint64
	Balance int64
	UserID  string
}
