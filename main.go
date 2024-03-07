package main

// @title           Health Data Service
// @version         1.0
// @description     This is a simple Health Data Service.

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  Health Data Service OpenAPI
// @externalDocs.url

import (
	"health-data-service/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run() // listen and serve on localhost:8080
}
