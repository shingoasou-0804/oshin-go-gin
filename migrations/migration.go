package main

import (
	"github.com/shingoasou-0804/oshin-go-gin/infra"
	"github.com/shingoasou-0804/oshin-go-gin/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}
