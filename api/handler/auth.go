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


// UpdateAuth godoc
// @Summary Update user profile
// @Description Update the user's authentication profile using their ID.
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body pb.User true "User data"
// @Success 202 {string} string "SUCCESS"
// @Failure 400 {object} gin.H "StatusBadRequest"
// @Failure 500 {object} gin.H "StatusInternalServerError"
// @Router /auths/profile/:id [put]
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


// DeleteAuth godoc
// @Summary Delete user profile
// @Description Update the user's authentication profile using their ID.
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 202 {string} string "SUCCESS"
// @Failure 400 {object} gin.H "StatusBadRequest"
// @Failure 500 {object} gin.H "StatusInternalServerError"
// @Router /auths/profile/:id [delete]
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


// ShowProfile godoc
// @Summary ShowProfiel user profile
// @Description Update the user's authentication profile using their ID.
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Failure 400 {object} gin.H "StatusBadRequest"
// @Failure 500 {object} gin.H "StatusInternalServerError"
// @Router /auths/profile/:id [get]
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


