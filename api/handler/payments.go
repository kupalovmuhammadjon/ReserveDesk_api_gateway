package handler

import (
	"api_gateway_service/genproto/payments"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) MakePayment(c *gin.Context) {
	req := &payments.Payment{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return	
	}
	resp, err := h.Payments.MakePayment(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	req := &payments.PaymentsFilter{Id: id}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Info("id bo`sh", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Payments.GetPayments(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	req := &payments.Payment{Id: id}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Info("id bo`sh", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Payments.UpdatePayment(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeletePayment(c *gin.Context) {
	id := c.Param("id")
	req := &payments.Id{Id: id}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Payments.DeletePayment(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetStatus(c *gin.Context) {
	id := c.Param("id")
	req := &payments.Id{Id: id}
	if err := c.ShouldBindQuery(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Payments.GetStatus(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
