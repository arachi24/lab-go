package config

import (
	"github.com/spf13/viper"
)

func NewConfig() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.ReadInConfig()

}
