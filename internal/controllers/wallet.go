package controllers

type WalletServiceI interface {
}

type WalletController struct {
	service WalletServiceI
}

func NewWalletController(service WalletServiceI) *WalletController {
	return &WalletController{service: service}
}
