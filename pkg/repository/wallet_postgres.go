package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/morpheus228/infotecs_golang_test/models"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (r *WalletPostgres) CreateWallet(wallet models.Wallet) error {
	query := fmt.Sprintf("INSERT INTO %s (id, balance) values ($1, $2)", walletsTable)
	_, err := r.db.Exec(query, wallet.Id, wallet.Balance)

	if err != nil {
		return err
	}

	return nil
}

func (r *WalletPostgres) MakeTransaction(transaction models.Transaction) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (amount, from_wallet_id, to_wallet_id, created_at) values ($1, $2, $3, $4) RETURNING id", transactionsTable)
	row := tx.QueryRow(query, transaction.Amount, transaction.FromWalletId, transaction.ToWalletId, transaction.CreatedAt)

	var transactionId int
	if err = row.Scan(&transactionId); err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET balance = balance - $1 WHERE id = $2", walletsTable)
	_, err = tx.Exec(query, transaction.Amount, transaction.FromWalletId)

	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET balance = balance + $1 WHERE id = $2", walletsTable)
	_, err = tx.Exec(query, transaction.Amount, transaction.ToWalletId)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *WalletPostgres) GetWalletHistory(walletId string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := fmt.Sprintf(`SELECT * FROM %s WHERE from_wallet_id = $1 OR to_wallet_id = $1`, transactionsTable)

	if err := r.db.Select(&transactions, query, walletId); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *WalletPostgres) GetWallet(walletId string) (models.Wallet, error) {
	var wallet models.Wallet
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, walletsTable)
	err := r.db.Get(&wallet, query, walletId)
	return wallet, err
}
