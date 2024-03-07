package dicom

import (
	"fmt"
	"health-data-service/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	GetAttributeParams struct {
		Tag string `form:"tag" binding:"required"`
	}
)

// @Summary      Extract DICOM header attribute
// @Description  Extracts a DICOM header attribute based on the provided tag.
// @Tags         dicom
// @Accept       json
// @Produce      json
// @Param 		 id 	path 	string					true   	"Identifier of the DICOM file"
// @Param        tag    query   GetAttributeParams   	true   	"Identifier of the uploaded DICOM file"
// @Success      200  {array}   dicom.Element
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/dicom/{id}/attribute [get]
func GET(s services.DicomService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileID := c.Param("id")

		var attributeParams GetAttributeParams
		if err := c.ShouldBind(&attributeParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("Bad Request - %v", err.Error()),
			})
			return
		}

		attribute, err := s.ExtractHeaderAttribute(fileID, &attributeParams.Tag)
		if err != nil {
			c.JSON(err.StatusCode(), gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, attribute)
	}
}

// @Summary      Convert DICOM file to a png image
// @Description  Convert DICOM file to a png image.
// @Tags         dicom
// @Accept       json
// @Produce      image/png
// @Param 		 id 	path 	string	true	"Identifier of the DICOM file"
// @Success      200  {string}  string  "PNG image content"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/dicom/{id}/image [get]
func ConvertImage(s services.DicomService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileID := c.Param("id")
		image, err := s.ConvertFileToImage(fileID)
		if err != nil {
			c.JSON(err.StatusCode(), gin.H{
				"error": fmt.Sprintf("Failed to convert image - %v", err.Error()),
			})
			return
		}

		c.Data(http.StatusOK, "image/png", image)
	}
}
