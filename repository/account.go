package account_repository

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	database "github.com/golu360/internal-transfers/db"
	"github.com/golu360/internal-transfers/db/models"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAccount(accountId string) (*models.Account, error) {
	db, err := database.GetDb()
	if err != nil {
		zap.L().Error("Could not connect to db", zap.Error(err), zap.Any("accountId", accountId))
		return nil, fiber.ErrInternalServerError
	}
	account := new(models.Account)
	err = db.First(&account, "account_id = ?", accountId).Select("account_id", "balance").Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		}
		return nil, fiber.ErrInternalServerError
	}
	return account, nil
}

func DebitFunds(accountId string, amount decimal.Decimal) error {
	db, err := database.GetDb()
	if err != nil {
		zap.L().Error("Could not connect to db", zap.Error(err), zap.Any("accountId", accountId))
		return fiber.ErrInternalServerError
	}
	results := db.Clauses(clause.Locking{ // will block other queries from updating or deleting the row
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).
		Model(&models.Account{}).Where("account_id = ?", accountId).Update("balance",
		gorm.Expr("balance - ?", amount))
	return results.Error
}

func CreditFunds(accountId string, amount decimal.Decimal) error {
	db, err := database.GetDb()
	if err != nil {
		zap.L().Error("Could not connect to db", zap.Error(err), zap.Any("accountId", accountId))
		return fiber.ErrInternalServerError
	}
	results := db.Clauses(clause.Locking{ // will block other queries from updating or deleting the row
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).
		Model(&models.Account{}).Where("account_id = ?", accountId).Update("balance",
		gorm.Expr("balance + ?", amount))
	return results.Error
}
