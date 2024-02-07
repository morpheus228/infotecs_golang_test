package models

type Transaction struct {
	Id           int     `json:"-" db:"id"`
	CreatedAt    string  `json:"time" db:"created_at"`
	Amount       float32 `json:"amount" binding:"required" db:"amount"`
	ToWalletId   string  `json:"to" binding:"required" db:"to_wallet_id"`
	FromWalletId string  `json:"from" db:"from_wallet_id"`
}
