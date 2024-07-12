package handler

import (
	"api_gateway_service/genproto/menu"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMenu(c *gin.Context) {
	req := &menu.MenuRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("ma'lumot kelmadi", "error: ", err.Error())
		return
	}
	resp, err := h.Menu.CreateMenu(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("ma'lumot kelmadi", "error: ", err.Error())
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) UpdateMenu(c *gin.Context) {
	req := &menu.MenuUpateRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("ma'lumot kelmadi", "error: ", err.Error())
		return
	}
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		h.Logger.Info("id bo`shqayti ")
		return
	}
	resp, err := h.Menu.UpdateMenu(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
func (h *Handler) DeleteMenu(c *gin.Context) {
	req := &menu.Id{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("ma'lumotni olishda xato", "err: ", err.Error())
		return
	}
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		h.Logger.Info("id bo`sh")
		return
	}
	resp, err := h.Menu.DeleteMenu(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)

}
func (h *Handler) GetByIdMenu(c *gin.Context) {
	req := &menu.Id{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = c.Param("id")
	if req.Id == "" {
		h.Logger.Info("id bo`sh")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}
	resp, err := h.Menu.GetByIdMenu(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
func (h *Handler) GetAllMenu(c *gin.Context) {
	req := &menu.MenuFilter{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = c.Param("id")
	if req.Id == "" {
		h.Logger.Error("id bo`sh")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}
	resp, err := h.Menu.GetAllMenu(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
