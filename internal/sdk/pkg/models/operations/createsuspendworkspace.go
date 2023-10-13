// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"singlestore/internal/sdk/pkg/models/shared"
)

type CreateSuspendWorkspaceRequest struct {
	// ID of the workspace
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

type CreateSuspendWorkspaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateWorkspace *shared.CreateWorkspace
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
