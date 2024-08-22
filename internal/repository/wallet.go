package repository

import (
	"github.com/jmoiron/sqlx"
)

type WalletRepo struct {
	db *sqlx.DB
}

func NewWalletRepo(db *sqlx.DB) *WalletRepo {
	return &WalletRepo{db: db}
}

func (r *WalletRepo) Exists(userID string) (bool, error) {
	q := `SELECT EXISTS (SELECT 1 FROM wallets WHERE user_id=$1)`
	var exists bool
	err := r.db.Get(&exists, q, userID)
	if err != nil {
		return false, err
	}
	return exists, nil
}
