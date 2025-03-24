package account_service

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/golu360/internal-transfers/db"
	"github.com/golu360/internal-transfers/dtos"
	repository "github.com/golu360/internal-transfers/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateAccount(body *dtos.CreateAccountDto) error {
	return repository.CreateAccount(body.AccountId, body.Balance)
}

func GetAccount(accountId int64) (*dtos.GetAccountResponseDto, error) {
	account, err := repository.GetAccount(accountId)
	if err != nil {
		return nil, err
	}
	return &dtos.GetAccountResponseDto{
		AccountId: account.AccountId,
		Balance:   account.Balance,
	}, nil
}

func TransferFunds(body *dtos.CreateTransactionDto) error {
	db, err := database.GetDb()
	if err != nil {
		zap.L().Error("Could not connect to db", zap.Error(err), zap.Any("body", body))
		return fiber.ErrInternalServerError
	}
	sourceAccount, err := repository.GetAccount(body.SourceAccountId)
	if err != nil {
		zap.L().Error("Error fetching source account", zap.Error(err), zap.Int64("sourceAccountId", body.SourceAccountId))
		return err
	}
	if sourceAccount.Balance.Cmp(body.Amount) == -1 { // check if enough balance in account
		zap.L().Error("insufficient balance in source acccount")
		return fiber.NewError(fiber.ErrBadRequest.Code, "Insufficient Balance")
	}

	// fetch destination account
	_, err = repository.GetAccount(body.DestinationAccountId)
	if err != nil {
		zap.L().Error("Error fetching destination account", zap.Error(err), zap.Int64("destinationAccountId", body.DestinationAccountId))
		return err
	}
	db.Transaction(func(tx *gorm.DB) error {
		if err = repository.CreateTransaction(body.SourceAccountId, body.DestinationAccountId, body.Amount); err != nil {
			return err
		}
		if err = repository.DebitFunds(body.SourceAccountId, body.Amount); err != nil {
			return err
		}

		if err = repository.CreditFunds(body.DestinationAccountId, body.Amount); err != nil {
			return err
		}
		return nil
	})

	return nil
}
