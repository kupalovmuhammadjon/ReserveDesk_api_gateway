package pkg

import (
	"api_gateway_service/config"
	pbuAuthservice "api_gateway_service/genproto/auth"
	pbuOrderservice "api_gateway_service/genproto/order"
	pbuReservation "api_gateway_service/genproto/reservations"
	payments "api_gateway_service/genproto/payments"
	restaurant "api_gateway_service/genproto/restaurant"
	menu "api_gateway_service/genproto/menu"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthenticationClient(cfg *config.Config) pbuAuthservice.AuthClient {
	conn, err := grpc.NewClient("localhost"+cfg.AUTH_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error while connecting authentication service ", err)
	}

	return pbuAuthservice.NewAuthClient(conn)
}

func NewOrderClient(cfg *config.Config) pbuOrderservice.OrderServiceClient {
	conn, err := grpc.NewClient("localhost"+cfg.RESERVATION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("erro while connecting order service ", err)
	}

	return pbuOrderservice.NewOrderServiceClient(conn)
}

func NewReservationClient(cfg *config.Config) pbuReservation.ReservationServiceClient {
	conn, err := grpc.NewClient("localhost"+cfg.RESERVATION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error while connecting reservation service ", err)
	}

	return pbuReservation.NewReservationServiceClient(conn)
}

func NewMenuClient(cfg *config.Config) menu.MenuServiceClient {
	conn, err := grpc.NewClient("localhost"+cfg.RESERVATION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error while connecting menu service ", err)
	}

	return menu.NewMenuServiceClient(conn)
}


func NewRestaurantClient(cfg *config.Config) restaurant.RestaurantClient {
	conn, err := grpc.NewClient("localhost"+cfg.RESERVATION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error while connecting Restaurant service ", err)
	}

	return restaurant.NewRestaurantClient(conn)
}

func NewPaymentsClient(cfg *config.Config) payments.PaymentsClient {
	conn, err := grpc.NewClient("localhost"+cfg.RESERVATION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error while connecting payments service ", err)
	}

	return payments.NewPaymentsClient(conn)
}
