package services_test

import (
	"fmt"
	"health-data-service/api/helper"
	"health-data-service/api/services"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const LocalTestFile = "../../testfiles/IM000001"
const TestPath = "files"

func TestUploadFile(t *testing.T) {
	var test = map[string]struct {
		filePath    string
		checkHeader bool
		header      multipart.FileHeader
		err         *helper.Error
	}{
		"When valid file and valid header Return uuid": {
			filePath:    LocalTestFile,
			checkHeader: false,
			header: multipart.FileHeader{
				Size: 3642242,
			},
			err: nil,
		},
		"When valid path and invalid header Return error": {
			filePath:    LocalTestFile,
			checkHeader: true,
			header: multipart.FileHeader{
				Size: 1,
			},
			err: &helper.Error{
				Code:    http.StatusBadRequest,
				Message: "Could not parse file - not enough bytes left until buffer limit to complete this operation",
			},
		},
		"When invalid path and invalid header Return error": {
			filePath:    "",
			checkHeader: true,
			header: multipart.FileHeader{
				Size: 1,
			},
			err: &helper.Error{
				Code:    http.StatusBadRequest,
				Message: "Could not parse file - invalid argument",
			},
		},
	}
	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			// SETUP
			mockRepo := new(MockRepository)
			service := services.NewService(mockRepo)
			var testFile *os.File

			var err error
			if tt.filePath != "" {
				testFile, err = os.Open(tt.filePath)
				if err != nil {
					t.Fatal(err)
				}
				defer testFile.Close()
			}

			type workAroundFile struct {
				file multipart.File
			}
			testFileMap := workAroundFile{
				file: testFile,
			}

			// TEST
			uuid, actualErr := service.UploadFile(testFileMap.file, tt.header)

			// ASSERTS
			if tt.err != nil && actualErr != nil {
				assert.Nil(t, uuid)
			} else {
				assert.NotEmpty(t, uuid)
			}
			assert.Equal(t, tt.err, actualErr)

			// Clean Directory
			if !tt.checkHeader && tt.filePath != "" {
				if removeErr := os.RemoveAll(TestPath); removeErr != nil {
					fmt.Println(removeErr)
				}
			}
		})
	}
}
