package main

import (

	"api_gateway_service/api"
	"api_gateway_service/config"
)

func main() {
	cfg := config.Config{}

	router := api.NewRouter(&cfg)
	router.Run(cfg.HTTP_PORT)

}
