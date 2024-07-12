package handler

import (
	"api_gateway_service/config"
	pbAuthservice "api_gateway_service/genproto/auth"
	menu "api_gateway_service/genproto/menu"
	pbOrderservice "api_gateway_service/genproto/order"
	payments "api_gateway_service/genproto/payments"
	pbReservation "api_gateway_service/genproto/reservations"
	restaurant "api_gateway_service/genproto/restaurant"
	"api_gateway_service/pkg"
)

type Handler struct {
	ClientAuthentication pbAuthservice.AuthClient
	ClientOrder          pbOrderservice.OrderServiceClient
	ClientReservation    pbReservation.ReservationServiceClient
	Menu                 menu.MenuServiceClient
	Payments             payments.PaymentsClient
	Restaurant           restaurant.RestaurantClient
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		ClientAuthentication: pkg.NewAuthenticationClient(cfg),
		ClientOrder:          pkg.NewOrderClient(cfg),
		ClientReservation:    pkg.NewReservationClient(cfg),
		//Menu: pkg,
		//Payments: pkg,
		//Restaurant: pkg,
	}
}
