package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	AccountId int64           `gorm:"index;primaryKey;not null" json:"accountId"`
	Balance   decimal.Decimal `gorm:"type:decimal(7,4);default:0.00" json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
