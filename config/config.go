package config

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

const (
	LogLevel    = "info"
	production  = "prod"
	develop     = "dev"
	SecurityKey = "SecretYouShouldHide"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	AppPort                string `mapstructure:"APP_PORT"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func (e *Env) IsProduction() bool {
	return e.AppEnv == production
}

func (e *Env) IsDevelop() bool {
	return e.AppEnv == develop
}

func (e *Env) GetPort() string {
	return ":" + e.AppPort
}

func (e *Env) GetDbUser() string {
	return e.DBUser
}

func (e *Env) GetDbPassword() string {
	return e.DBPass
}

func (e *Env) GetDbHost() string {
	return e.DBHost
}

func (e *Env) GetDbPort() string {
	return e.DBPort
}

func (e *Env) GetDbSchema() string {
	return e.DBName
}

func GetJWTSigningMethod() jwt.SigningMethod {
	return jwt.SigningMethodHS256
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
