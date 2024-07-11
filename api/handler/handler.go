package handler

import (
	"api_service/config"
	pbuAuthservice "api_service/genproto/auth"
	pbuOrderservice "api_service/genproto/order"
	pbuReservation "api_service/genproto/reservations"
	"api_service/pkg"
)

type Handler struct {
	ClientAuthentication pbuAuthservice.AuthClient
	ClientOrder          pbuOrderservice.OrderServiceClient
	ClientReservation    pbuReservation.ReservationServiceClient
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		ClientAuthentication: pkg.NewAuthenticationClient(cfg),
		ClientOrder: pkg.NewOrderClient(cfg),
		ClientReservation: pkg.NewReservationClient(cfg),
	}
}
