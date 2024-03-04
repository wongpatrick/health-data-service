package routes

import (
	"health-data-service/api/handler/dicom"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	//  routes
	// v1.GET("/dicom", dicom.GET)
	v1.POST("/dicom/upload", dicom.POST)
}
