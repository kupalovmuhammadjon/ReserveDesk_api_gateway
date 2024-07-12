package handler

import (
	pb "api_gateway_service/genproto/order"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateOrder(c *gin.Context) {
	order := pb.Order{}

	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":   err.Error(),
			"Message": "Error while decoding",
		})
		log.Println("Error while decoding")
		return
	}
	tctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err = h.ClientOrder.CreateOrder(tctx, &order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while creating invitation",
		})
		log.Println("Error while creating invitation ", err)
		return
	}

	c.JSON(http.StatusCreated, "SUCCESS")
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}

	orderUpdate := pb.Updateorder{}

	err := json.NewDecoder(c.Request.Body).Decode(&orderUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL body: %s", err.Error())
		return
	}

	orderUpdate.Id = id

	tctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err = h.ClientOrder.UpdateOrder(tctx, &orderUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, "SUCCESS")

}

func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL body: %s", err.Error())
		return
	}

	tcxt, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err := h.ClientOrder.DeleteOrder(tcxt, &pb.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, "SUCCESS")
}

func (h *Handler) GetByIdOrder(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL : %s", err.Error())
		return
	}

	tcxt, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	orderInfo, err := h.ClientOrder.GetOrderById(tcxt, &pb.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, orderInfo)
}

func (h *Handler) GetAllOrder(c *gin.Context) {
	orderfilter := pb.OrderFilter{}

	err := json.NewDecoder(c.Request.Body).Decode(&orderfilter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL body: %s", err.Error())
		return
	}
	tctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	orders, err := h.ClientOrder.GetAllOrder(tctx, &orderfilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, orders)
}
