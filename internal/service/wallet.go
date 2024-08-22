package service

type WalletRepositoryI interface {
	Exists(userID string) (bool, error)
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
