package controlers

import (
	"Murtazo/app/structs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var shablon structs.User

	fmt.Printf("shablon: %v\n", shablon)
}


func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://192.168.43.45:5500")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	}

	c.Next()
}
