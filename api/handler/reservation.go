package handler

import (
	pb "api_gateway_service/genproto/reservations"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateReservation(c *gin.Context) {
	reservation := pb.Reservation{}

	err := json.NewDecoder(c.Request.Body).Decode(&reservation)
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

	_, err = h.ClientReservation.CreateReservation(tctx, &reservation)
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

func (h *Handler) UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}

	reservation := pb.ReservationUpdate{}

	if err := json.NewDecoder(c.Request.Body).Decode(&reservation); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":   err.Error(),
			"Message": "Error while decoding",
		})
		log.Println("Error while decoding")
		return
	}

	if _, err := uuid.Parse(reservation.UserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}

	reservation.Id = id

	tctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err := h.ClientReservation.UpdateReservation(tctx, &reservation)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while creating invitation",
		})
		log.Println("Error while creating invitation ", err)
		return
	}

	c.JSON(201, "SUCCESS")
}

func (h *Handler) DeleteReservation(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}

	tctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err := h.ClientReservation.DeleteReservation(tctx, &pb.Id{Id: id})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while creating invitation",
		})
		log.Println("Error while creating invitation ", err)
		return
	}

	c.JSON(201, "SUCCESS")

}

func (h *Handler) GetByIdReservation(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}

	tctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	reservation, err := h.ClientReservation.GetReservationById(tctx, &pb.Id{Id: id})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while creating invitation",
		})
		log.Println("Error while creating invitation ", err)
		return
	}

	c.JSON(201, reservation)
}

func (h *Handler) GetAllReservation(c *gin.Context) {
	reservationF := pb.ReservationFilter{}

	err := json.NewDecoder(c.Request.Body).Decode(&reservationF)
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

	reservations, err := h.ClientReservation.GetAllReservations(tctx, &reservationF)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, reservations)
}