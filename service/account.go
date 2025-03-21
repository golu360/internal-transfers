package account_service

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	database "github.com/golu360/internal-transfers/db"
	"github.com/golu360/internal-transfers/db/models"
	"github.com/golu360/internal-transfers/dtos"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateAccount(body *dtos.CreateAccountDto) error {
	db, err := database.GetDb()
	if err != nil {
		zap.L().Error("Could not connect to db", zap.Error(err), zap.Any("request", body))
		return fiber.ErrInternalServerError
	}
	account := &models.Account{
		AccountId: body.AccountId,
		Balance:   body.Balance,
	}
	result := db.Create(account)
	if err := result.Error; err != nil {
		zap.L().Error("Error occurred while trying to insert record", zap.Error(err))
		if database.IsDuplicateKeyError(err) {
			return fiber.ErrConflict
		}
		return fiber.ErrInternalServerError
	}
	zap.L().Info("Account created successfully", zap.String("accountId", body.AccountId.String()))
	return nil
}

func GetAccount(accountId string) (*dtos.GetAccountResponseDto, error) {
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
	return &dtos.GetAccountResponseDto{
		AccountId: account.AccountId,
		Balance:   account.Balance,
	}, nil

}
