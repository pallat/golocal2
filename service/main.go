package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const delay time.Duration = 200 * time.Millisecond

func main() {
	r := gin.Default()
	r.GET("/:number", func(c *gin.Context) {

		time.Sleep(delay)

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"number":  c.Param("number"),
		})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
