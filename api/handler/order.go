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


// @Summary Create order 
// @Description Create the order's authentication order body.
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body order.Order true "Order data"
// @Success 202 {string} string "SUCCESS"
// @Failure 400 {object} gin.H "StatusBadRequest"
// @Failure 500 {object} gin.H "StatusInternalServerError"
// @Router /orders [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	order := pb.Order{}

	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
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
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while creating invitation",
		})
		log.Println("Error while creating invitation ", err)
		return
	}

	c.JSON(http.StatusCreated, "SUCCESS")
}


// @Summary Update order 
// @Description Update the order's authentication order body.
// @Tags order
// @Accept  json
// @Produce  json
// @Param  id path string true "User ID"
// @Param order body order.Updateorder true "Order data"
// @Success 202 {string} string "SUCCESS"
// @Failure 400 {object} gin.H "StatusBadRequest"
// @Failure 500 {object} gin.H "StatusInternalServerError"
// @Router /orders/:id [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		h.Logger.Info("Error with getting Id from URL ", "err: ", err.Error())
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
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
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
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, "SUCCESS")

}


// @Summary Delete order 
// @Description Delete the order's authentication order body.
// @Tags order
// @Accept  json
// @Produce  json
// @Param  id path string true "User ID"
// @Success 202 {string} string "SUCCESS"
// @Failure 400 {object} gin.H "StatusBadRequest"
// @Failure 500 {object} gin.H "StatusInternalServerError"
// @Router /orders/:id [delete]
func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		h.Logger.Info("id bo`sh", "err: ", err.Error())
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
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, "SUCCESS")
}

// @Summary Get By Id order 
// @Description Update the order's authentication order body.
// @Tags order
// @Accept  json
// @Produce  json
// @Param  id path string true "User ID"
// @Success 202 {string} string "SUCCESS"
// @Failure 400 {object} gin.H "StatusBadRequest"
// @Failure 500 {object} gin.H "StatusInternalServerError"
// @Router /orders/:id [get]
func (h *Handler) GetByIdOrder(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		h.Logger.Error("id bo`sh", "err: ", err.Error())
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
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, orderInfo)
}


// @Summary Get By All order 
// @Description Get by all the order's authentication order body.
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body order.OrderFilter true "Order data"
// @Success 202 {string} string "SUCCESS"
// @Failure 400 {object} gin.H "StatusBadRequest"
// @Failure 500 {object} gin.H "StatusInternalServerError"
// @Router /orders/:id [get]
func (h *Handler) GetAllOrder(c *gin.Context) {
	orderfilter := pb.OrderFilter{}

	err := json.NewDecoder(c.Request.Body).Decode(&orderfilter)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
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
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, orders)
}
