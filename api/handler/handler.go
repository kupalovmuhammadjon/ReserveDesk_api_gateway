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
	logger "api_gateway_service/pkg/logger"
	"log"
	"log/slog"
)

type Handler struct {
	ClientAuthentication pbAuthservice.AuthClient
	ClientOrder          pbOrderservice.OrderServiceClient
	ClientReservation    pbReservation.ReservationServiceClient
	Menu                 menu.MenuServiceClient
	Payments             payments.PaymentsClient
	Restaurant           restaurant.RestaurantClient
	Logger               *slog.Logger
}

func NewHandler(cfg *config.Config) *Handler {
	l, err := logger.New()
	if err != nil {
		log.Fatal("error: ", err)
	}
	return &Handler{
		ClientAuthentication: pkg.NewAuthenticationClient(cfg),
		ClientOrder:          pkg.NewOrderClient(cfg),
		ClientReservation:    pkg.NewReservationClient(cfg),
		Menu:                 pkg.NewMenuClient(cfg),
		Payments:             pkg.NewPaymentsClient(cfg),
		Restaurant:           pkg.NewRestaurantClient(cfg),
		Logger:               l,
	}
}
