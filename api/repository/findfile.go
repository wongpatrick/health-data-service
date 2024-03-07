package repository

import (
	"fmt"
	"health-data-service/api/helper"
	"net/http"

	"github.com/suyashkumar/dicom"
)

const localPath = "./files/dicom"

func (r *repository) FindFile(id string) (*dicom.Dataset, *helper.Error) {
	filePath := localPath + "/" + id + ".dcm"
	dicomData, err := dicom.ParseFile(filePath, nil)
	if err != nil {
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Could not parse file - %v", err.Error()),
		}
	}
	return &dicomData, nil
}
