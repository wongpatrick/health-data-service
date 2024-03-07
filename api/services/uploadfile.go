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

func (d *dicomService) UploadFile(file multipart.File, header multipart.FileHeader) (*string, *helper.Error) {
	defer file.Close()
	// To validate it is a dicom file
	if _, err := dicom.Parse(file, header.Size, nil); err != nil {
		return nil, &helper.Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Could not parse file - %v", err.Error()),
		}
	}
	// Returning back to beginning of the file
	file.Seek(0, io.SeekStart)

	if err := os.MkdirAll(localPath, os.ModePerm); err != nil {
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Could not create directory - %v", err.Error()),
		}
	}

	uuid, err := fileCreation(localPath, &file)
	if err != nil {
		return nil, err
	}

	return uuid, nil
}

func fileCreation(path string, file *multipart.File) (*string, *helper.Error) {
	uuid := helper.GenerateUUID()
	fullPath := path + "/" + uuid + ".dcm"

	destination, err := os.Create(fullPath)
	if err != nil {
		return nil, &helper.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if _, err := io.Copy(destination, *file); err != nil {
		destination.Close() // Did not use defer closed here since I cannot delete file if it's not closed
		if removeErr := os.Remove(destination.Name()); removeErr != nil {
			return nil, &helper.Error{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("File could not be saved & removed: %v - %v", err.Error(), removeErr.Error())}
		}
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("File could not be saved & removed: %v", err.Error())}
	}
	destination.Close()
	return &uuid, nil
}
