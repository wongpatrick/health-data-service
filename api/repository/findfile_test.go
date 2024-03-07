package repository_test

import (
	"health-data-service/api/helper"
	"health-data-service/api/repository"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFile(t *testing.T) {
	var test = map[string]struct {
		id  string
		err *helper.Error
	}{
		"When valid Id Return data": {
			id:  "test-file",
			err: nil,
		},
		"When valid Id and cannot parse file Return err": {
			id: "test",
			err: &helper.Error{
				Code:    http.StatusInternalServerError,
				Message: "Could not parse file - open files\\dicom\\test.dcm: The system cannot find the file specified.",
			},
		},
	}
	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			repo := repository.NewRepository()
			dataSet, err := repo.FindFile(tt.id)

			if tt.err != nil {
				assert.Nil(t, dataSet)
			} else {
				assert.NotEmpty(t, dataSet)
			}
			assert.Equal(t, tt.err, err)
		})
	}
}
