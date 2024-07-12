package pkg

import (
	"api_gateway_service/config"
	pbuAuthservice "api_gateway_service/genproto/auth"
	pbuOrderservice "api_gateway_service/genproto/order"
	pbuReservation "api_gateway_service/genproto/reservations"
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
	conn, err := grpc.NewClient("localhost"+cfg.ORDER_SERVICE_PORT,
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
