package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() (*gorm.DB, error) {
	dbHost := viper.GetString("db.host")
	dbUser := viper.GetString("db.user")
	dbPassword := viper.GetString("db.password")
	dbName := viper.GetString("db.name")
	dbPort := viper.GetString("db.port")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
