package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/morpheus228/infotecs_golang_test/models"
)

type Wallet interface {
	CreateWallet(wallet models.Wallet) error
	MakeTransaction(transaction models.Transaction) error
	GetWalletHistory(walletId string) ([]models.Transaction, error)
	GetWallet(walletId string) (models.Wallet, error)
}

type Repository struct {
	Wallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Wallet: NewWalletPostgres(db),
	}
}
