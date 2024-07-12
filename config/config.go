package config

import (
	"log"
	"os"
	"github.com/spf13/cast"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP_PORT           string
	API_GATEWAY_PORT    string
	AUTH_SERVICE_PORT   string
	ORDER_SERVICE_PORT  string
	RESERVATION_SERVICE string
	SIGNING_KEY         string
}

func Login() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load .env ", err)
	}

	cfg := Config{}

	cfg.API_GATEWAY_PORT = cast.ToString(coalesce("HTTP_PORT", ":8080"))
	cfg.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", ":7777"))
	cfg.ORDER_SERVICE_PORT = cast.ToString(coalesce("ORDER_SERVICE_PORT", ":8888"))
	cfg.RESERVATION_SERVICE = cast.ToString(coalesce("RESERVATION_SERVICE", ":9999"))
	cfg.SIGNING_KEY = cast.ToString(coalesce("SIGNING_KEY", "GARD"))

	return &cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
