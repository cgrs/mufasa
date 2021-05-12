package main

import (
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName(".mufasa")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")
	viper.SetDefault("max_retries", 3)
	viper.SetDefault("retry_timeout", time.Second*30)
	viper.SetDefault("totp_secret", "")
	viper.SetDefault("profile_name", "mfa")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}
