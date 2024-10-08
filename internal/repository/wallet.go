package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github/usmonzodasomon/wallet/internal/models"
)

type WalletRepo struct {
	db *sqlx.DB
}

func NewWalletRepo(db *sqlx.DB) *WalletRepo {
	return &WalletRepo{db: db}
}

func (r *WalletRepo) GetBalance(userID string) (int64, error) {
	q := `SELECT balance FROM wallets WHERE user_id=$1`
	var balance int64
	err := r.db.Get(&balance, q, userID)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (r *WalletRepo) AddBalance(walletID int64, amount int64) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	q := `UPDATE wallets SET balance = balance + $1 WHERE id=$2`
	_, err = tx.Exec(q, amount, walletID)
	if err != nil {
		return tx.Rollback()
	}

	q = `INSERT INTO transactions (wallet_id, amount) VALUES ($1, $2)`
	_, err = tx.Exec(q, walletID, amount)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (r *WalletRepo) GetWallet(userID string) (models.Wallet, error) {
	q := `SELECT id, balance, user_id, is_identified FROM wallets WHERE user_id=$1`
	var wallet models.Wallet
	err := r.db.Get(&wallet, q, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Wallet{}, models.ErrWalletNotFound
		}
		return models.Wallet{}, err
	}
	return wallet, nil
}

func (r *WalletRepo) TotalDeposits(walletID int64) (int64, int64, error) {
	q := `SELECT
    COUNT(*) AS total_count,
    SUM(amount) AS total_sum
FROM
    transactions
WHERE
    DATE_TRUNC('month', time) = DATE_TRUNC('month', CURRENT_DATE)
AND wallet_id=$1;
`

	var total models.TotalDeposits
	err := r.db.Get(&total, q, walletID)
	if err != nil {
		return 0, 0, err
	}
	return total.TotalCount, total.TotalSum, nil
}
