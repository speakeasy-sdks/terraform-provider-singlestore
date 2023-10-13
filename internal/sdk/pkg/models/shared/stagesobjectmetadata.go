// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type StagesObjectMetadataContentType string

const (
	StagesObjectMetadataContentTypeStr                         StagesObjectMetadataContentType = "str"
	StagesObjectMetadataContentTypeArrayOfStagesObjectMetadata StagesObjectMetadataContentType = "arrayOfStagesObjectMetadata"
)

type StagesObjectMetadataContent struct {
	Str                         *string
	ArrayOfStagesObjectMetadata []StagesObjectMetadata

	Type StagesObjectMetadataContentType
}

func CreateStagesObjectMetadataContentStr(str string) StagesObjectMetadataContent {
	typ := StagesObjectMetadataContentTypeStr

	return StagesObjectMetadataContent{
		Str:  &str,
		Type: typ,
	}
}

func CreateStagesObjectMetadataContentArrayOfStagesObjectMetadata(arrayOfStagesObjectMetadata []StagesObjectMetadata) StagesObjectMetadataContent {
	typ := StagesObjectMetadataContentTypeArrayOfStagesObjectMetadata

	return StagesObjectMetadataContent{
		ArrayOfStagesObjectMetadata: arrayOfStagesObjectMetadata,
		Type:                        typ,
	}
}

func (u *StagesObjectMetadataContent) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = StagesObjectMetadataContentTypeStr
		return nil
	}

	arrayOfStagesObjectMetadata := []StagesObjectMetadata{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfStagesObjectMetadata); err == nil {
		u.ArrayOfStagesObjectMetadata = arrayOfStagesObjectMetadata
		u.Type = StagesObjectMetadataContentTypeArrayOfStagesObjectMetadata
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u StagesObjectMetadataContent) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfStagesObjectMetadata != nil {
		return json.Marshal(u.ArrayOfStagesObjectMetadata)
	}

	return nil, nil
}

// StagesObjectMetadataFormat - Format of the response
type StagesObjectMetadataFormat string

const (
	StagesObjectMetadataFormatJSON                   StagesObjectMetadataFormat = "json"
	StagesObjectMetadataFormatLessThanNilGreaterThan StagesObjectMetadataFormat = "<nil>"
)

func (e StagesObjectMetadataFormat) ToPointer() *StagesObjectMetadataFormat {
	return &e
}

func (e *StagesObjectMetadataFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "<nil>":
		*e = StagesObjectMetadataFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StagesObjectMetadataFormat: %v", v)
	}
}

// StagesObjectMetadataType - Object type
type StagesObjectMetadataType string

const (
	StagesObjectMetadataTypeUnknown                StagesObjectMetadataType = ""
	StagesObjectMetadataTypeJSON                   StagesObjectMetadataType = "json"
	StagesObjectMetadataTypeDirectory              StagesObjectMetadataType = "directory"
	StagesObjectMetadataTypeLessThanNilGreaterThan StagesObjectMetadataType = "<nil>"
)

func (e StagesObjectMetadataType) ToPointer() *StagesObjectMetadataType {
	return &e
}

func (e *StagesObjectMetadataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "json":
		fallthrough
	case "directory":
		fallthrough
	case "<nil>":
		*e = StagesObjectMetadataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StagesObjectMetadataType: %v", v)
	}
}

// StagesObjectMetadata - Represents the metadata corresponding to an object in Stages
type StagesObjectMetadata struct {
	Content *StagesObjectMetadataContent `json:"content,omitempty"`
	Created *string                      `json:"created,omitempty"`
	// Format of the response
	Format       *StagesObjectMetadataFormat `json:"format,omitempty"`
	LastModified *string                     `json:"last_modified,omitempty"`
	Mimetype     *string                     `json:"mimetype,omitempty"`
	// Name of the Stages object
	Name *string `json:"name,omitempty"`
	// Path of the Stages object
	Path *string `json:"path,omitempty"`
	Size *int64  `json:"size,omitempty"`
	// Object type
	Type     *StagesObjectMetadataType `json:"type,omitempty"`
	Writable *bool                     `json:"writable,omitempty"`
}