package main

import (
	"github.com/gin-gonic/gin"
)

var orders = map[string]interface{}{
	"order001": map[string]interface{}{
		"item":     "Laptop",
		"quantity": 1,
		"price":    1200.50,
	},
	"order002": map[string]interface{}{
		"item":     "Mouse",
		"quantity": 2,
		"price":    25.99,
	},
}

func main() {
	r := gin.Default()

	r.GET("/order", func(c *gin.Context) {
		c.JSON(200, gin.H{"messge": "oder service", "orders": orders})
	})
	r.Run(":8081")
}
