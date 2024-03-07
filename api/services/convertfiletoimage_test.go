package services_test

import (
	"health-data-service/api/helper"
	"health-data-service/api/services"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suyashkumar/dicom"
)

func TestConvertFileToImage(t *testing.T) {
	var test = map[string]struct {
		id          string
		isPixelData bool
		err         *helper.Error
	}{
		"When valid id Return image data": {
			id:          "test-file",
			isPixelData: true,
			err:         nil,
		},
		"When valid id with no pixel data return error": {
			id:          "test-file",
			isPixelData: false,
			err: &helper.Error{
				Code:    http.StatusInternalServerError,
				Message: "Could not find pixel data - element not found",
			},
		},
	}
	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			// SETUP
			mockRepo := new(MockRepository)
			var dicomData dicom.Dataset
			var err error
			if tt.isPixelData {
				dicomData, err = dicom.ParseFile(LocalTestFile, nil)
				if err != nil {
					t.Fatal(err)
				}
			}

			expectedError := &helper.Error{}
			expectedError = nil
			mockRepo.On("FindFile", tt.id).Return(&dicomData, expectedError)

			service := services.NewService(mockRepo)

			// TEST
			element, err := service.ConvertFileToImage(tt.id)

			// ASSERTS
			if tt.err != nil {
				assert.Nil(t, element)
			} else {
				assert.NotEmpty(t, element) // In an ideal state, I would be comparing files
			}
			assert.Equal(t, tt.err, err)

		})
	}
}
