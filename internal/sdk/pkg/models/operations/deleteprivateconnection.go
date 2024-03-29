// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"singlestore/internal/sdk/pkg/models/shared"
)

type DeletePrivateConnectionRequest struct {
	// ID of the private connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionID"`
}

type DeletePrivateConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	DeletePrivateConnection *shared.DeletePrivateConnection
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
