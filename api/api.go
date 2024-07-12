package api

import (
	"api_gateway_service/api/handler"
	"api_gateway_service/api/middleware"
	"api_gateway_service/config"

	"github.com/gin-gonic/gin"
	_ "api_gateway_service/api/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ReserveDesk API
// @version 1.0
// @description ReserveDesk is program to book seats in restaurants order food before arrival.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
// @BasePath /reservedesk.uz

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http



func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/reservedesk.uz")
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
	auth.PUT("/profile/:id", h.UpdateAuth)
	auth.DELETE("/profile/:id", h.DeleteAuth)
	auth.GET("/profile/:id", h.ShowProfile)

	menu := api.Group("/menu")
	menu.POST("/menu", h.CreateMenu)
	menu.PUT("/menu/:id", h.UpdateMenu)
	menu.DELETE("/menu/:id", h.DeleteMenu)
	menu.GET("/get/menu/:id", h.GetByIdMenu)
	menu.GET("/getAll/menu/:id", h.GetAllMenu)

	payment := api.Group("/payment")
	payment.POST("/make/payment", h.MakePayment)
	payment.GET("/get/payment/:id", h.GetPayment)
	payment.PUT("/payment/:id", h.UpdatePayment)
	payment.DELETE("/payment/:id", h.DeletePayment)
	payment.GET("/get/status/:id", h.GetStatus)

	restaurant := api.Group("/restaurant")
	restaurant.POST("/restaurant/:id", h.CreateRestaurant)
	restaurant.GET("/restaurant/:id", h.GetRestaurantById)
	restaurant.PUT("/restaurant/:id", h.UpdateRestaurantById)
	restaurant.DELETE("/restaurant/:id", h.DeleteRestaurantById)

	return router
}
