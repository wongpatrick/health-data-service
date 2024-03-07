package routes

import (
	"health-data-service/api/handler/dicom"
	"health-data-service/api/repository"
	"health-data-service/api/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	repo := repository.NewRepository()
	dicomService := services.NewService(repo)

	// DICOM routes
	v1.GET("/dicom/:id/attribute", dicom.GET(dicomService))
	v1.GET("/dicom/:id/image", dicom.ConvertImage(dicomService))
	v1.POST("/dicom", dicom.POST(dicomService))
}
