package handler

import (
	"api_gateway_service/genproto/payments"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// MakePayment
// @Summary Create Payments
// @Description This API for Creating Menu
// @Tags payments
// @Accept json
// @Produce json
// @Param Payments body payments.Payment true "Payments"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /make/payment [post]
func (h *Handler) MakePayment(c *gin.Context) {
	req := &payments.Payment{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
		return	
	}
	resp, err := h.Payments.MakePayment(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetPayment
// @Summary Get Payments
// @Description This API for get wth id Menu
// @Tags payments
// @Accept json
// @Produce json
// @Param Payments body payments.PaymentsFilter true "Payments"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /get/payment/:id [get]
func (h *Handler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	req := &payments.PaymentsFilter{Id: id}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Info("id bo`sh", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	resp, err := h.Payments.GetPayments(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdatePayment
// @Summary Update Payments
// @Description This API for updating Menu
// @Tags payments
// @Accept json
// @Produce json
// @Param Payments body payments.Payment true "Payments"
// @Param   id     path    string     true        "Payment ID"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /payment/{id} [put]
func (h *Handler) UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	req := &payments.Payment{Id: id}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Info("id bo`sh", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	resp, err := h.Payments.UpdatePayment(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}

// DeletePayment
// @Summary Delete Payments
// @Description This API for get wth id Menu
// @Tags payments
// @Accept json
// @Produce json
// @Param   id     path    string     true        "Payment ID"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /payment/:id [delete]
func (h *Handler) DeletePayment(c *gin.Context) {
	id := c.Param("id")
	req := &payments.Id{Id: id}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}	// _ "api_gateway_service/api/docs"

	resp, err := h.Payments.DeletePayment(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}

// GetStatus
// @Summary Get Payments
// @Description This API for get wth id Menu
// @Tags payments
// @Accept json
// @Produce json
// @Param   id     path    string     true        "Payment ID"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /get/status/:id [get]
func (h *Handler) GetStatus(c *gin.Context) {
	id := c.Param("id")
	req := &payments.Id{Id: id}
	if err := c.ShouldBindQuery(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	resp, err := h.Payments.GetStatus(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}
