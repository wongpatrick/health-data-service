package services

import (
	"health-data-service/api/helper"
	repo "health-data-service/api/repository"
	"mime/multipart"

	"github.com/suyashkumar/dicom"
)

type DicomService interface {
	ConvertFileToImage(id string, fileType string) ([]byte, *helper.Error)
	ExtractHeaderAttribute(id string, tagQuery *string) (*dicom.Element, *helper.Error)
	UploadFile(file multipart.File, header multipart.FileHeader) (*string, *helper.Error)
}

type dicomService struct {
	repository repo.Repository
}

func NewService(repo repo.Repository) DicomService {
	return &dicomService{
		repository: repo,
	}
}
