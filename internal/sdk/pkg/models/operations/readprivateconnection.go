// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"singlestore/internal/sdk/pkg/models/shared"
)

type ReadPrivateConnectionRequest struct {
	// ID of the private connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionID"`
	// Comma-separated values list that correspond to the filtered fields for returned entities
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
}

type ReadPrivateConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PrivateConnections []shared.PrivateConnection
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
