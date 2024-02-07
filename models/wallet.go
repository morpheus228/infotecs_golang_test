package models

type Wallet struct {
	Id      string  `json:"id" db:"id"`
	Balance float32 `json:"balance" binding:"required" db:"balance"`
}
