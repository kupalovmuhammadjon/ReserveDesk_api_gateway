package main

import (
	"api_gateway_service/api"
	"api_gateway_service/config"
)

func main() {
	cfg := config.Load()

	router := api.NewRouter(cfg)
	router.Run(cfg.HTTP_PORT)
}
