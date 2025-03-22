package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	TransactionId        uuid.UUID       `gorm:"type:uuid;index;primaryKey;not null" json:"transactionId"`
	SourceAccountId      uuid.UUID       `gorm:"type:uuid;not null" json:"sourceAccountId"`
	DestinationAccountId uuid.UUID       `gorm:"type:uuid;not null" json:"destinationAccountId"`
	Amount               decimal.Decimal `gorm:"type:decimal(7,4);default:0.00" json:"amount"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
