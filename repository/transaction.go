package repository

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/golu360/internal-transfers/db"
	"github.com/golu360/internal-transfers/db/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func CreateTransaction(sourceAccountId int64, destinationAccountId int64, amount decimal.Decimal) error {
	db, err := database.GetDb()
	if err != nil {
		zap.L().Error("Could not connect to db", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	transactionId := uuid.New()
	transaction := &models.Transaction{
		TransactionId:        transactionId,
		SourceAccountId:      sourceAccountId,
		DestinationAccountId: destinationAccountId,
		Amount:               amount,
	}
	result := db.Create(transaction)
	if err := result.Error; err != nil {
		zap.L().Error("Error occurred while trying to create transaction", zap.Error(err))
		if database.IsDuplicateKeyError(err) {
			return fiber.ErrConflict
		}
		return fiber.ErrInternalServerError
	}
	zap.L().Info("Created transaction successfull", zap.Any("transaction", transaction))
	return nil
}
