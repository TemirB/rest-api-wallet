package models

import "github.com/google/uuid"

type Wallet struct {
	Id      uuid.UUID `json:"walletId"` // ID of the wallet
	Balance float64   // Balance of the wallet
}

type WalletRequest struct {
	Id            uuid.UUID `json:"walletId"`      // ID of the wallet !(в pdf указано поле "valletId")
	OperationType string    `json:"operationType"` // Type of operation
	Amount        float64   `json:"amount"`        // Balance of the wallet
}
