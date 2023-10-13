// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// StorageDRSetup - Represents the information specified to setup Storage DR
type StorageDRSetup struct {
	// List of database names
	DatabaseNames []string `json:"databaseNames"`
	// Region ID of the secondary region
	RegionID string `json:"regionID"`
}