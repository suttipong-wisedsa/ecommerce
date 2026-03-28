package main

import (
	"myapi/config"
	"myapi/handlers"
	"myapi/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/products", handlers.GetProduct)
	r.GET("/products/cart", handlers.GetCart)
	r.POST("/products", handlers.AddToCart)

	r.Run(":8080")
}
