package main

import (
	"time"

	"github.com/239103/alice_agent/models"
	"github.com/gin-gonic/gin"
)

func actionEndpoint(c *gin.Context) {

}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			// Create response struct of APIClient type
			var res models.APIClient
			res.Type = "ping"
			res.Status = "success"
			res.Message.ID = uint32(time.Now().Unix())
			res.Message.Data = "pong"
			res.Message.TimeStamp = time.Now()

			c.JSON(200, &res)
		})
		v1.POST("/action", actionEndpoint)
	}

	// gin.SetMode(gin.ReleaseMode)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
