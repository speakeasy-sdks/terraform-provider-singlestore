// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"singlestore/internal/sdk/pkg/models/shared"
)

type ListWorkspaceGroupRequest struct {
	// Comma-separated values list that correspond to the filtered fields for returned entities
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// To include any terminated workspace groups, set to `true`
	IncludeTerminated *bool `queryParam:"style=form,explode=true,name=includeTerminated"`
}

type ListWorkspaceGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WorkspaceGroups []shared.WorkspaceGroup
}
