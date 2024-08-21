package service

type WalletRepositoryI interface {
}

type WalletService struct {
	repo WalletRepositoryI
}

func NewWalletService(repo WalletRepositoryI) *WalletService {
	return &WalletService{repo: repo}
}
