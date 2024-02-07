package service

import (
	"math/rand"
	"time"

	"github.com/morpheus228/infotecs_golang_test/models"
	"github.com/morpheus228/infotecs_golang_test/pkg/repository"
)

const startWalletBalance = 100.0

type WalletService struct {
	repo repository.Wallet
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) GenerateWalletId() string {
	rand.Seed(time.Now().UnixNano())
	characters := "abcdefghijklmnopqrstuvwxyz0123456789"
	idLength := 32
	id := make([]byte, idLength)

	for i := range id {
		id[i] = characters[rand.Intn(len(characters))]
	}

	return string(id)
}

func (s *WalletService) CreateWallet() (models.Wallet, error) {
	wallet := models.Wallet{Id: s.GenerateWalletId(), Balance: startWalletBalance}

	if err := s.repo.CreateWallet(wallet); err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (s *WalletService) GetWallet(walletId string) (models.Wallet, error) {
	return s.repo.GetWallet(walletId)
}

func (s *WalletService) GetWalletHistory(walletId string) ([]models.Transaction, error) {
	return s.repo.GetWalletHistory(walletId)
}

func (s *WalletService) MakeTransaction(fromWalletId string, transaction models.Transaction) error {
	transaction.FromWalletId = fromWalletId
	transaction.CreatedAt = time.Now().Format(time.RFC3339)
	return s.repo.MakeTransaction(transaction)
}
