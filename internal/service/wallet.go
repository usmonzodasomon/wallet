package service

import "github/usmonzodasomon/wallet/internal/models"

type WalletRepositoryI interface {
	Exists(userID string) (bool, error)
	GetBalance(userID string) (int64, error)
}

type WalletService struct {
	repo WalletRepositoryI
}

func NewWalletService(repo WalletRepositoryI) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) Exists(userID string) (bool, error) {
	return s.repo.Exists(userID)
}

func (s *WalletService) GetBalance(userID string) (float64, error) {
	exists, err := s.Exists(userID)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, models.ErrWalletNotFound
	}

	balanceInt, err := s.repo.GetBalance(userID)
	if err != nil {
		return 0, err
	}

	balance := float64(balanceInt) / 100
	return balance, nil
}
