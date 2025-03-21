package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
}

func main() {
}
