package handler

import (
	"api_service/config"
	pbuAuthservice "api_service/genproto/auth"
	pbuOrderservice "api_service/genproto/order"
	pbuReservation "api_service/genproto/reservations"
	"api_service/pkg"
	pbuMenu "api_gateway/genproto/menu"
	pbuPayments "api_gateway/genproto/payments"
	pbuRestaurant "api_gateway/genproto/restaurant"
)

type Handler struct {
	ClientAuthentication pbuAuthservice.AuthClient
	ClientOrder          pbuOrderservice.OrderServiceClient
	ClientReservation    pbuReservation.ReservationServiceClient
	ClientRestaurant     pbuRestaurant.RestaurantClient
	ClientMenu           pbuMenu.MenuServiceClient
	ClienstPayment       pbuPayments.PaymentsClient
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		ClientAuthentication: pkg.NewAuthenticationClient(cfg),
		ClientOrder: pkg.NewOrderClient(cfg),
		ClientReservation: pkg.NewReservationClient(cfg),


// func NewHendler(R restaurant.RestaurantClient, M menu.MenuServiceClient, P payments.PaymentsClient) *Handler {
// 	return &Handler{
// 		Restaurant: R,
// 		Menu:       M,
// 		Payments:   P,
	}
}
