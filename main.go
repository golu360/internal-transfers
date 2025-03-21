package main

import (
	"github.com/gofiber/fiber/v2"
	endpoints "github.com/golu360/internal-transfers/constants"
	database "github.com/golu360/internal-transfers/db"
	"github.com/golu360/internal-transfers/db/models"
	"github.com/golu360/internal-transfers/dtos"
	account_service "github.com/golu360/internal-transfers/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	db, err := database.GetDb()
	if err != nil {
		zap.L().Panic("Error occurred while trying to auto migrate", zap.Error(err))
	}
	zap.L().Debug("Migrating accounts schema")
	db.AutoMigrate(&models.Account{})
	database.Close(db)
}

func main() {
	app := fiber.New()

	app.Get(endpoints.HEALTH_CHECK, func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post(endpoints.CREATE_ACCOUNT, func(c *fiber.Ctx) error {
		body := new(dtos.CreateAccountDto)
		if err := c.BodyParser(body); err != nil {
			return fiber.ErrInternalServerError
		}
		if err := account_service.CreateAccount(body); err != nil {
			return err
		}
		return c.SendStatus(201)
	})

	app.Listen(":" + viper.GetString("app.port"))
}
