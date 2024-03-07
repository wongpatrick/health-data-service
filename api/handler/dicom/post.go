package dicom

import (
	"fmt"
	"health-data-service/api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	DicomUploadResponse struct {
		Id string `json:"id"`
	}
)

// @Summary      Upload a dicom file
// @Description  Allows the user to upload a dicom file and returns a unique identifier
// @Tags         dicom
// @Accept       multipart/form-data
// @Produce      json
// @Param		 file formData file true "DICOM file to upload"
// @Success      200  {object}  DicomUploadResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/dicom [post]
func POST(s services.DicomService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("POST v1/dicom")
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Failed to upload file - %v", err.Error()),
			})
			return
		}
		defer file.Close()

		identifier, storeErr := s.UploadFile(file, *header)
		if storeErr != nil {
			c.JSON(storeErr.StatusCode(), gin.H{
				"error": fmt.Sprintf("Failed to save file - %v", storeErr.Error()),
			})
			return
		}
		// TODO: Figure out the format for return
		c.JSON(http.StatusCreated, DicomUploadResponse{Id: *identifier})
	}

}
