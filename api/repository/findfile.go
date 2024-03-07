package repository

import (
	"fmt"
	"health-data-service/api/helper"
	"net/http"
	"path/filepath"

	"github.com/suyashkumar/dicom"
)

func (r *repository) FindFile(id string) (*dicom.Dataset, *helper.Error) {
	filePath := filepath.Join(helper.LocalPath, id+".dcm")
	dicomData, err := dicom.ParseFile(filePath, nil)
	if err != nil {
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Could not parse file - %v", err.Error()),
		}
	}
	return &dicomData, nil
}
