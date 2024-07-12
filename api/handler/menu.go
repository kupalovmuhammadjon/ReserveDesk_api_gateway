package handler

import (
	"api_gateway_service/genproto/menu"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMenu
// @Summary Create Menu
// @Description This API for creating Menu
// @Tags menu
// @Accept json
// @Produce json
// @Param Menu body menu.MenuRequest true "Menu"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /menu [post]
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


// UpdateMenu
// @Summary Update Menu
// @Description This API for updating Menu
// @Tags menu
// @Accept json
// @Produce json
// @Param Menu body menu.MenuUpateRequest true "Menu"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /menu/{id} [put]
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

// DeleteMenu
// @Summary Delete Menu
// @Description This API for deleting Menu
// @Tags menu
// @Accept json
// @Produce json
// @Param   id     path    string     true        "Menu ID"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /menu/{id} [delete]
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

// GetByIdMenu
// @Summary Get Menu
// @Description This API for get wth id Menu
// @Tags menu
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param   id     path    string     true        "Menu ID"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /menu/get/menu/{id} [get]

func (h *Handler) GetByIdMenu(c *gin.Context) {
	req := &menu.Id{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	req.Id = c.Param("id")
	if req.Id == "" {
		h.Logger.Info("id bo`sh")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
	}
	resp, err := h.Menu.GetByIdMenu(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllMenu
// @Summary GetAll Menu
// @Description This API for get wth id Menu
// @Tags menu
// @Accept json
// @Produce json
// @Param Menu query menu.MenuFilter true "Menu"
// @Param   id     path    string     true        "Restaurant ID"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /getAll/menu/{id} [get]
func (h *Handler) GetAllMenu(c *gin.Context) {
	req := &menu.MenuFilter{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	req.Id = c.Param("id")
	if req.Id == "" {
		h.Logger.Error("id bo`sh")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
	}
	resp, err := h.Menu.GetAllMenu(c, req)
	if err != nil {
		h.Logger.Error("ma'lumot kelmadi", "err: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err: %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}
