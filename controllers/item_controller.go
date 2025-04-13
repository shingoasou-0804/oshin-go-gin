package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shingoasou-0804/oshin-go-gin/services"
)

type IItemController interface {
	FindAll(ctx *gin.Context)
}

type ItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

func (c *ItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
