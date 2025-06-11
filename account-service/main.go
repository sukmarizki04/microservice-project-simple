package main
import (
	"github.com/gin-gonic/gin"
)
var accounts = map[string]interface{}{
	"user01": map[string]interface{}{
		"name": "Alice",
		"email": "alice@gmail.com",
	},
	"user011": map[string]interface{}{
		"name": "Alice",
		"email": "alice@gmail.com",
	},
}


func main(){
	r := gin.Default()
	r.GET("/account", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "account service", "accounts": accounts})
	})
	r.Run(":8082")
}