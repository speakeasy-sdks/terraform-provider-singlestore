// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ReplicatedDatabaseDuplicationState - Duplication state of the database
type ReplicatedDatabaseDuplicationState string

const (
	ReplicatedDatabaseDuplicationStatePending  ReplicatedDatabaseDuplicationState = "Pending"
	ReplicatedDatabaseDuplicationStateActive   ReplicatedDatabaseDuplicationState = "Active"
	ReplicatedDatabaseDuplicationStateInactive ReplicatedDatabaseDuplicationState = "Inactive"
	ReplicatedDatabaseDuplicationStateError    ReplicatedDatabaseDuplicationState = "Error"
)

func (e ReplicatedDatabaseDuplicationState) ToPointer() *ReplicatedDatabaseDuplicationState {
	return &e
}

func (e *ReplicatedDatabaseDuplicationState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Pending":
		fallthrough
	case "Active":
		fallthrough
	case "Inactive":
		fallthrough
	case "Error":
		*e = ReplicatedDatabaseDuplicationState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReplicatedDatabaseDuplicationState: %v", v)
	}
}

// ReplicatedDatabase - Represents information related to a database's replication status
type ReplicatedDatabase struct {
	// Name of the database
	DatabaseName string `json:"databaseName"`
	// Duplication state of the database
	DuplicationState ReplicatedDatabaseDuplicationState `json:"duplicationState"`
	// Name of the region
	Region string `json:"region"`
}
