package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -X POST  localhost:3001/centrifugo/subscribe -v
func main() {
	r := gin.Default()
	r.POST("/centrifugo/subscribe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": make(map[string]int),
		})
	})

	r.Run(":3001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
