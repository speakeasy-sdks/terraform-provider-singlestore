// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"singlestore/internal/sdk/pkg/models/shared"
)

type ReadStagesFileRequest struct {
	// If enabled, the API request returns only metadata
	Metadata *bool `queryParam:"style=form,explode=true,name=metadata"`
	// Path in Stages to a file or folder
	Path string `pathParam:"style=simple,explode=false,name=path"`
	// ID of the Stages-enabled workspace group
	WorkspaceGroupID string `pathParam:"style=simple,explode=false,name=workspaceGroupID"`
}

type ReadStagesFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	StagesObjectMetadata *shared.StagesObjectMetadata
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}