package api

import (
	"api_service/api/handler"
	"api_service/api/middleware"
	"api_service/config"

	"github.com/gin-gonic/gin"
)

// @title Auth Service API
// @version 1.0
// @description This is a sample server for Auth Service.
// @host localhost:7777
// @schemes http
func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	api := router.Group("/reservation")
	api.Use(middleware.JWTMiddleware())

	h := handler.NewHandler(cfg)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	auth.PUT("/profile/:id", h.UpdateAuth)
	auth.DELETE("/profile/:id", h.DeleteAuth)
	auth.GET("/profile/:id", h.ShowProfile)

	return router
}
