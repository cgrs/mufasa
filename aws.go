package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/pquerna/otp/totp"
)

var cfg *aws.Config
var client *sts.Client

func loadConfig() *aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}

func getCallerId() string {
	id, err := client.GetCallerIdentity(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return *id.Arn
}

func getToken(serial, otp string) (*sts.GetSessionTokenOutput, error) {
	return client.GetSessionToken(context.TODO(), &sts.GetSessionTokenInput{
		SerialNumber: &serial,
		TokenCode:    &otp,
	})
}

func generateToken(serial string, totpSecret string, maxRetries int, timeout time.Duration) *OutputCredentials {
	tries := 0
	if serial == "" {
		callerId := getCallerId()
		serial = strings.Replace(callerId, "user", "mfa", 1)
	}
	var res *sts.GetSessionTokenOutput
	for tries < maxRetries {
		otp, err := totp.GenerateCode(totpSecret, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		res, err = getToken(serial, otp)
		if err != nil {
			log.Printf("error while retrieving token, retrying in %d seconds (%d of %d)...\n", int(timeout.Seconds()), tries, maxRetries)
			log.Println(err)
			time.Sleep(timeout)
			tries++
		} else {
			break
		}
	}
	if res == nil {
		log.Printf("error while retrieving token (max retries)\n")
		return nil
	}
	return FromCredentials(res.Credentials)
}

func init() {
	cfg = loadConfig()
	client = sts.NewFromConfig(*cfg)
}
