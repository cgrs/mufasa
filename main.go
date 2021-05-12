package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	serial := viper.GetString("arn")
	totp := viper.GetString("totp_secret")
	if totp == "" {
		log.Fatal("please provide totp secret to generate OTP")
	}
	maxRetries := viper.GetInt("max_retries")
	timeout := viper.GetDuration("retry_timeout")
	prefix := viper.GetString("profile_name")
	out := generateToken(serial, totp, maxRetries, timeout)
	if out != nil {
		fmt.Println(out.ToINI(prefix))
	}
}
