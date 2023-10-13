// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"singlestore/internal/sdk/pkg/models/shared"
)

type CreateStagesFileRequest struct {
	CreateStagesFileRequest *shared.CreateStagesFileRequest `request:"mediaType=multipart/form-data"`
	// If set to `true`, forces creation of an empty file
	IsFile *bool `queryParam:"style=form,explode=true,name=isFile"`
	// Path in Stages
	Path string `pathParam:"style=simple,explode=false,name=path"`
	// ID of the Stages-enabled workspace group
	WorkspaceGroupID string `pathParam:"style=simple,explode=false,name=workspaceGroupID"`
}

type CreateStagesFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateStagesFile *shared.CreateStagesFile
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
