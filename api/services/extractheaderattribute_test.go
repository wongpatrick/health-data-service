package services_test

import (
	"health-data-service/api/helper"
	"health-data-service/api/services"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/suyashkumar/dicom"
)

// Mock Repository should be in it's own file
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindFile(id string) (*dicom.Dataset, *helper.Error) {
	args := m.Called(id)

	// Check if the arguments provided to the method are valid
	if dataset, ok := args.Get(0).(*dicom.Dataset); ok {
		return dataset, args.Get(1).(*helper.Error)
	}

	// Return default values if the arguments are not as expected
	return nil, nil
}

func TestExtractHeaderAttribute(t *testing.T) {
	var test = map[string]struct {
		id       string
		tagQuery *string
		err      *helper.Error
	}{
		"When valid id and valid tagQuery Return data": {
			id:       "test-file",
			tagQuery: helper.ToStringPtr("(0002,0001)"),
			err:      nil,
		},
		"When valid id and invalid tagQuery Return data": {
			id:       "test-file",
			tagQuery: helper.ToStringPtr("test,0001"),
			err: &helper.Error{
				Code:    http.StatusInternalServerError,
				Message: "Group could not be parsed - strconv.ParseUint: parsing \"test\": invalid syntax",
			},
		},
		"When valid id and valid tagQuery but can't find element Return data": {
			id:       "test-file",
			tagQuery: helper.ToStringPtr("0001,0001"),
			err: &helper.Error{
				Code:    http.StatusBadRequest,
				Message: "element not found",
			},
		},
	}
	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			// SETUP
			mockRepo := new(MockRepository)
			dicomData, err := dicom.ParseFile(LocalTestFile, nil)
			if err != nil {
				t.Fatal(err)
			}

			expectedError := &helper.Error{}
			expectedError = nil
			mockRepo.On("FindFile", tt.id).Return(&dicomData, expectedError)

			service := services.NewService(mockRepo)

			// TEST
			element, err := service.ExtractHeaderAttribute(tt.id, tt.tagQuery)

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
