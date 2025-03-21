package dtos

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateAccountDto struct {
	AccountId uuid.UUID       `json:"account_id"`
	Balance   decimal.Decimal `json:"initial_balance"`
}

type GetAccountResponseDto struct {
	AccountId uuid.UUID       `json:"account_id"`
	Balance   decimal.Decimal `json:"balance"`
}

type CreateTransactionDto struct {
	SourceAccountId      uuid.UUID       `json:"source_account_id"`
	DestinationAccountId uuid.UUID       `json:"destination_account_id"`
	Amount               decimal.Decimal `json:"amount"`
}
