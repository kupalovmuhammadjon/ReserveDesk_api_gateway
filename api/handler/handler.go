package handler

import (
	menu "api_gateway/genproto/menu"
	payments "api_gateway/genproto/payments"
	restaurant "api_gateway/genproto/restaurant"
)

type Handler struct {
	Restaurant restaurant.RestaurantClient
	Menu       menu.MenuServiceClient
	Payments   payments.PaymentsClient
}

func NewHendler(R restaurant.RestaurantClient, M menu.MenuServiceClient, P payments.PaymentsClient) *Handler {
	return &Handler{
		Restaurant: R,
		Menu:       M,
		Payments:   P,
	}
}
