package api

import (
	"api_gateway/api/handler"
	"api_gateway/config"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	//api.Use(middleware.JWTMiddleware())
	h := handler.Handler{}

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

	return r
}
