package repository

import (
	"health-data-service/api/helper"

	"github.com/suyashkumar/dicom"
)

type Repository interface {
	FindFile(id string) (*dicom.Dataset, *helper.Error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
