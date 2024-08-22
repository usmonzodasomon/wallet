package service

import "github/usmonzodasomon/wallet/internal/models"

type WalletRepositoryI interface {
	GetBalance(userID string) (int64, error)
	AddBalance(walletID int64, amount int64) error
	GetWallet(userID string) (models.Wallet, error)
	TotalDeposits(walletID int64) (int64, int64, error)
}

type WalletService struct {
	repo WalletRepositoryI
}

func NewWalletService(repo WalletRepositoryI) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) Exists(userID string) (bool, error) {
	_, err := s.repo.GetWallet(userID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *WalletService) GetBalance(userID string) (float64, error) {
	_, err := s.Exists(userID)
	if err != nil {
		return 0, err
	}

	balanceInt, err := s.repo.GetBalance(userID)
	if err != nil {
		return 0, err
	}

	balance := float64(balanceInt) / 100
	return balance, nil
}

const (
	MaxBalanceUnidentified = 10_000 * 100
	MaxBalanceIdentified   = 100_000 * 100
)

func (s *WalletService) AddBalance(userID string, amount int64) error {
	wallet, err := s.repo.GetWallet(userID)
	if err != nil {
		return err
	}

	var MaxBalance int64
	if !wallet.IsIdentified {
		MaxBalance = MaxBalanceUnidentified * 100
	} else {
		MaxBalance = MaxBalanceIdentified * 100
	}

	if wallet.Balance+amount > MaxBalance {
		return models.ErrMaxBalanceExceeded
	}

	err = s.repo.AddBalance(wallet.ID, amount)
	if err != nil {
		return err
	}

	return nil
}

func (s *WalletService) TotalDeposits(userID string) (int64, float64, error) {
	wallet, err := s.repo.GetWallet(userID)
	if err != nil {
		return 0, 0, err
	}

	totalCount, totalSum, err := s.repo.TotalDeposits(wallet.ID)
	if err != nil {
		return 0, 0, err
	}
	return totalCount, float64(totalSum) / 100, nil
}
