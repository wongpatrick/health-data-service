package services

import (
	"fmt"
	"health-data-service/api/helper"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/suyashkumar/dicom"
)

const localPath = "./files/dicom"

func StoreFile(file multipart.File, header multipart.FileHeader) (*string, *helper.Error) {
	defer file.Close()
	if _, err := dicom.Parse(file, header.Size, nil); err != nil {
		return nil, &helper.Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Could not parse file - %v", err.Error()),
		}
	}
	file.Seek(0, io.SeekStart)

	if err := os.MkdirAll(localPath, os.ModePerm); err != nil {
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	uuid := helper.GenerateUUID()
	path := localPath + "/" + uuid + ".dcm"

	destination, err := os.Create(path)
	if err != nil {
		return nil, &helper.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	defer destination.Close()

	//TODO:
	// 1. Figure out naming convention

	if _, err := io.Copy(destination, file); err != nil {
		if removeErr := os.Remove(destination.Name()); removeErr != nil {
			return nil, &helper.Error{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("File could not be saved & removed: %v - %v", err.Error(), removeErr.Error())}
		}
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("File could not be saved & removed: %v", err.Error())}
	}

	return &uuid, nil
}
