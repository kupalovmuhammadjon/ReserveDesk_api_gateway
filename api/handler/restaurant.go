package handler

import (
	"api_gateway_service/genproto/restaurant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateRestaurant(c *gin.Context) {
	req := &restaurant.RestaurantCreate{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Restaurant.CreateRestaurant(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetRestaurantById(c *gin.Context) {
	id := c.Param("id")

	req := &restaurant.RestaurantFilter{Id: id}
	if err := c.ShouldBind(req); err != nil {
		h.Logger.Info("id bo`sh", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Restaurant.GetRestaurants(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateRestaurantById(c *gin.Context) {
	id := c.Param("id")
	req := &restaurant.RestaurantUpdate{Id: id}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Restaurant.UpdateRestaurant(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteRestaurantById(c *gin.Context) {
	id := c.Param("id")
	req := &restaurant.Id{Id: id}
	if err := c.ShouldBind(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Restaurant.DeleteRestaurant(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
