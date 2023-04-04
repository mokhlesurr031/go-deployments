package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type JWTToken struct {
	ExpiredIn time.Duration
	MaxAge    int
	Secret    string
}

var (
	jwt JWTToken
)

func JWT() *JWTToken {
	return &jwt
}

func loadToken() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	jwt = JWTToken{
		ExpiredIn: viper.GetDuration("jwt.TOKEN_EXPIRED_IN"),
		MaxAge:    viper.GetInt("jwt.TOKEN_MAXAGE"),
		Secret:    viper.GetString("jwt.TOKEN_SECRET"),
	}
}
