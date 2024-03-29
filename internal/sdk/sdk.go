// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"fmt"
	"net/http"
	"singlestore/internal/sdk/pkg/models/shared"
	"singlestore/internal/sdk/pkg/utils"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.singlestore.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Singlestore - SingleStore Management API: The `Management` API can be used to create and manage workspaces, workspace groups, private connections, etc. SingleStore recommends reading the [`Management` API Overview](https://docs.singlestore.com/managed-service/en/reference/management-api.html) before getting started with the API reference.
//
// All the URLs referenced in this API documentation use the `https://api.singlestore.com` service endpoint as their base.
type Singlestore struct {
	// Billing - Operations related to billing
	Billing *billing
	// Organizations - Operations related to organizations
	Organizations      *organizations
	PrivateConnections *privateConnections
	// Regions - Operations related to regions
	Regions *regions
	// Stages - Operations related to stages
	Stages          *stages
	WorkspaceGroups *workspaceGroups
	// Workspaces - Operations related to workspaces
	Workspaces *workspaces

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Singlestore)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Singlestore) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Singlestore) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Singlestore) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Singlestore) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Singlestore) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Singlestore {
	sdk := &Singlestore{
		sdkConfiguration: sdkConfiguration{
			Language:          "terraform",
			OpenAPIDocVersion: "1.1.33",
			SDKVersion:        "0.1.3",
			GenVersion:        "2.154.1",
			UserAgent:         "speakeasy-sdk/terraform 0.1.3 2.154.1 1.1.33 singlestore",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Billing = newBilling(sdk.sdkConfiguration)

	sdk.Organizations = newOrganizations(sdk.sdkConfiguration)

	sdk.PrivateConnections = newPrivateConnections(sdk.sdkConfiguration)

	sdk.Regions = newRegions(sdk.sdkConfiguration)

	sdk.Stages = newStages(sdk.sdkConfiguration)

	sdk.WorkspaceGroups = newWorkspaceGroups(sdk.sdkConfiguration)

	sdk.Workspaces = newWorkspaces(sdk.sdkConfiguration)

	return sdk
}
