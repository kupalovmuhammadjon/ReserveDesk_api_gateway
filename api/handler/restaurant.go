package handler

import (
	"api_gateway_service/genproto/restaurant"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateRestaurant
// @Summary Create Restaurant
// @Description This API for creating Menu
// @Tags restaurant
// @Accept json
// @Produce json
// @Param Restaurant body restaurant.RestaurantCreate true "Restaurant"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /restaurant/:id [post]
func (h *Handler) CreateRestaurant(c *gin.Context) {
	req := &restaurant.RestaurantCreate{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	resp, err := h.Restaurant.CreateRestaurant(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusCreated, resp)
}

// GetRestaurantById
// @Summary Get Restaurant
// @Description This API for get wth id Menu
// @Tags restaurant
// @Accept json
// @Produce json
// @Param   id     path    string     true        "Restaurant ID"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /restaurant/:id [get]
func (h *Handler) GetRestaurantById(c *gin.Context) {
	id := c.Param("id")

	req := &restaurant.RestaurantFilter{Id: id}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	resp, err := h.Restaurant.GetRestaurants(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateRestaurantById
// @Summary Update Payments
// @Description This API for updating Menu
// @Tags restaurant
// @Accept json
// @Produce json
// @Param   id     path    string     true        "Restaurant ID"
// @Param Payments body restaurant.RestaurantUpdate true "Payments"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /restaurant/:id [put]
func (h *Handler) UpdateRestaurantById(c *gin.Context) {
	id := c.Param("id")
	req := &restaurant.RestaurantUpdate{Id: id}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	resp, err := h.Restaurant.UpdateRestaurant(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteRestaurantById
// @Summary Delete Restaurant
// @Description This API for updating Menu
// @Tags restaurant
// @Accept json
// @Produce json
// @Param   id     path    string     true        "Restaurant ID"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /restaurant/:id [delete]
func (h *Handler) DeleteRestaurantById(c *gin.Context) {
	id := c.Param("id")
	req := &restaurant.Id{Id: id}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	resp, err := h.Restaurant.DeleteRestaurant(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}
