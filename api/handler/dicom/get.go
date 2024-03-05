package dicom

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	DicomAtrributesParams struct {
		Tag *string `form:"tag"`
	}
)

// @Summary      Extract DICOM header attribute
// @Description  Extracts a DICOM header attribute based on the provided tag.
// @Tags         dicom
// @Accept       json
// @Produce      json
// @Param 		 id 	path 	string 					true   	"Identifier of the DICOM file"
// @Param        tag    query   DicomAtrributesParams   true   	"Identifier of the uploaded DICOM file"
// @Success      200  {array}   map[string]string
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/dicom/{id} [get]
func GET(c *gin.Context) {
	log.Printf("GET DICOM HEADER")
	fileID := c.Param("id")

	c.JSON(http.StatusOK, map[string]string{fileID: fileID})
}
