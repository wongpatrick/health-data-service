package services

import (
	"bytes"
	"fmt"
	"health-data-service/api/helper"
	"image"
	"image/png"
	"net/http"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func (d *dicomService) ConvertFileToImage(id string) ([]byte, *helper.Error) {
	dicomData, err := d.repository.FindFile(id)
	if err != nil {
		return nil, err
	}

	pixelDataElement, elementErr := dicomData.FindElementByTag(tag.PixelData)
	if elementErr != nil {
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Could not find pixel data - %v", elementErr.Error()),
		}
	}

	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelDataElement.Value)

	images := make([]image.Image, len(pixelDataInfo.Frames))
	for i, fr := range pixelDataInfo.Frames {
		img, err := fr.GetImage()
		if err != nil {
			return nil, &helper.Error{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("Could not convert to image - %v", err.Error()),
			}
		}
		images[i] = img
	}

	var buffer bytes.Buffer
	if err := png.Encode(&buffer, images[0]); err != nil {
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Could not convert to png - %v", err.Error()),
		}
	}

	return buffer.Bytes(), nil
}
