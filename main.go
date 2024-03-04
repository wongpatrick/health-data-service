package main

// @title           Health Data Service
// @version         1.0
// @description     This is a simple Health Data Service.

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  Health Data Service OpenAPI
// @externalDocs.url

import (
	"log"
	"time"

	"health-data-service/api/routes"

	"github.com/gin-gonic/gin"
)

// A very basic Logger for the sake of the assignment
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {

	r := gin.Default()
	r.Use(Logger())
	routes.SetupRoutes(r)

	r.Run() // listen and serve on localhost:8080
}
