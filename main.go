package main

import (
	"github.com/golu360/internal-transfers/db"
	"github.com/golu360/internal-transfers/db/models"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	db, err := db.GetDb()
	if err != nil {
		zap.L().Panic("Error occurred while trying to auto migrate", zap.Error(err))
	}
	zap.L().Debug("Migrating accounts schema")
	db.AutoMigrate(&models.Account{})
}

func main() {
}
