// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Organization - Represents information related to an organization
type Organization struct {
	// The list of allowed IP addresses which can access the Management API
	FirewallRanges []string `json:"firewallRanges,omitempty"`
	// Name of the organization
	Name *string `json:"name,omitempty"`
	// ID of the organization
	OrgID string `json:"orgID"`
}
