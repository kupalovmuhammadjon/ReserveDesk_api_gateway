package handler

import (
	menu "api_gateway_service/genproto/menu"
	payments "api_gateway_service/genproto/payments"
	restaurant "api_gateway_service/genproto/restaurant"
)

type Handler struct {
	Restaurant restaurant.RestaurantClient
	Menu       menu.MenuServiceClient
	Payments   payments.PaymentsClient
}

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
