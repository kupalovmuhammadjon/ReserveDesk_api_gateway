package handler

import (
	pb "api_gateway_service/genproto/auth"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) UpdateAuth(c *gin.Context) {
	id := c.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}
	userp := pb.User{}

	err = json.NewDecoder(c.Request.Body).Decode(&userp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL body: %s", err.Error())
		return
	}
	userp.Id = id

	nestedctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.ClientAuthentication.UpdateProfile(nestedctx, &userp)
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

func (h *Handler) DeleteAuth(c *gin.Context) {
	id := c.Param("id")
	nestedctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if _,err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		return 
	}


	_, err := h.ClientAuthentication.DeleteProfile(nestedctx, &pb.Id{Id: id})
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

func (h *Handler) ShowProfile(c *gin.Context) {
	id := c.Param("id")

	nestedctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		return 
	}

	userp, err := h.ClientAuthentication.ShowProfile(nestedctx, &pb.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, userp)
}


