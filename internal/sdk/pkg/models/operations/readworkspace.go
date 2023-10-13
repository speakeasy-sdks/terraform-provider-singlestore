// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"singlestore/internal/sdk/pkg/models/shared"
)

type ReadWorkspaceRequest struct {
	// Comma-separated values list that correspond to the filtered fields for returned entities
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the workspace
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

type ReadWorkspaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Workspace *shared.Workspace
}
