package service

import (
	"github.com/morpheus228/infotecs_golang_test/models"
	"github.com/morpheus228/infotecs_golang_test/pkg/repository"
)

type Wallet interface {
	CreateWallet() (models.Wallet, error)
	GetWallet(walletId string) (models.Wallet, error)
	GetWalletHistory(walletId string) ([]models.Transaction, error)
	MakeTransaction(fromWalletId string, transaction models.Transaction) error
}

type Service struct {
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Wallet: NewWalletService(repos.Wallet),
	}
}
