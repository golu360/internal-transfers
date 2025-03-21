package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	AccountId uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"accountId"`
	Balance   decimal.Decimal `gorm:"type:decimal(7,6);" json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
