package dicom

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST dicom
// @Summary      Upload a dicom file
// @Description  Allows the user to upload a dicom file and returns a unique identifier
// @Tags         dicom
// @Accept       multipart/form-data
// @Produce      json
// @Param		 dicomFile formData file true "DICOM file to upload"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/dicom/upload [post]
func POST(c *gin.Context /*, uploader dicomFile.FileUploader*/) {
	log.Printf("POST v1/dicom/upload")
	_, _, err := c.Request.FormFile("dicomFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to read file - %v", err.Error()),
		})
		return
	}
	log.Printf("FOUND FILE")

	// created, postErr := services.PostTask(taskToCreate)
	// if postErr != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": postErr.Error(),
	// 	})
	// 	return
	// }
	c.JSON(http.StatusCreated, map[string]string{})
}
