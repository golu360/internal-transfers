package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	AccountId uuid.UUID       `gorm:"type:uuid;index;primaryKey;not null" json:"accountId"`
	Balance   decimal.Decimal `gorm:"type:decimal(7,6);default:0.00" json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
