package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shingoasou-0804/oshin-go-gin/controllers"
	"github.com/shingoasou-0804/oshin-go-gin/models"
	"github.com/shingoasou-0804/oshin-go-gin/repositories"
	"github.com/shingoasou-0804/oshin-go-gin/services"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}
	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)
	r := gin.Default()

	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.Run(":8080")
}
