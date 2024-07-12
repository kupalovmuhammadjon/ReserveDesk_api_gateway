package api

import (
	"api_service/api/handler"
	"api_service/api/middleware"
	"api_service/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	api := router.Group("/reservation")
	api.Use(middleware.JWTMiddleware())

	h := handler.NewHandler(cfg)

	// reservation api
	reservation := api.Group("/reservations")
	reservation.POST("/", h.CreateReservation)
	reservation.DELETE("/:id", h.DeleteReservation)
	reservation.PUT("/:id", h.UpdateReservation)
	reservation.GET("/", h.GetAllReservation)
	reservation.GET("/:id", h.GetByIdReservation)

	// order api
	order := api.Group("/orders")
	order.POST("/", h.CreateOrder)
	order.PUT("/:id", h.UpdateOrder)
	order.DELETE("/:id", h.DeleteAuth)
	order.GET("/", h.GetAllOrder)
	order.GET("/:id", h.GetByIdOrder)

	// auth api
	auth := api.Group("/auths")
	auth.PUT("/:id/profile", h.UpdateAuth)
	auth.DELETE("/:id/profile", h.DeleteAuth)
	auth.GET("/:id/profile", h.ShowProfile)

	return router
}
