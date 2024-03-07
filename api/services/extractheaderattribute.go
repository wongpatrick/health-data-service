package services

import (
	"fmt"
	"health-data-service/api/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func (d *dicomService) ExtractHeaderAttribute(id string, tagQuery *string) (*dicom.Element, *helper.Error) {
	parsedTag, parseErr := parseTag(tagQuery)
	if parseErr != nil {
		return nil, parseErr
	}

	dicomDataSet, err := d.repository.FindFile(id)
	if err != nil {
		return nil, err
	}

	element, parsedErr := dicomDataSet.FindElementByTag(*parsedTag)
	if parsedErr != nil {
		return nil, &helper.Error{
			Code:    http.StatusBadRequest,
			Message: parsedErr.Error(),
		}
	}

	return element, nil
}

func parseTag(tagString *string) (*tag.Tag, *helper.Error) {
	tagSplit := strings.Split(strings.Trim(*tagString, "()"), ",")

	group, err := strconv.ParseUint(tagSplit[0], 16, 16)
	if err != nil {
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Group could not be parsed - %v", err.Error()),
		}
	}

	element, err := strconv.ParseUint(tagSplit[1], 16, 16)
	if err != nil {
		return nil, &helper.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Element could not be parsed - %v", err.Error()),
		}
	}

	tagStruct := tag.Tag{
		Group:   uint16(group),
		Element: uint16(element),
	}

	return &tagStruct, nil
}
