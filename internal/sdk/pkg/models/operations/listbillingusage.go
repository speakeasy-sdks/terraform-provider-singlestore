// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"singlestore/internal/sdk/pkg/models/shared"
)

// ListBillingUsageAggregateBy - The aggregate type used to group usage which includes hour, day and month. default is hour
type ListBillingUsageAggregateBy string

const (
	ListBillingUsageAggregateByHour  ListBillingUsageAggregateBy = "hour"
	ListBillingUsageAggregateByDay   ListBillingUsageAggregateBy = "day"
	ListBillingUsageAggregateByMonth ListBillingUsageAggregateBy = "month"
)

func (e ListBillingUsageAggregateBy) ToPointer() *ListBillingUsageAggregateBy {
	return &e
}

func (e *ListBillingUsageAggregateBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "hour":
		fallthrough
	case "day":
		fallthrough
	case "month":
		*e = ListBillingUsageAggregateBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListBillingUsageAggregateBy: %v", v)
	}
}

// ListBillingUsageMetric - Metrics include ComputeCredit, StorageAvgByte. default is all
type ListBillingUsageMetric string

const (
	ListBillingUsageMetricComputeCredit  ListBillingUsageMetric = "ComputeCredit"
	ListBillingUsageMetricStorageAvgByte ListBillingUsageMetric = "StorageAvgByte"
)

func (e ListBillingUsageMetric) ToPointer() *ListBillingUsageMetric {
	return &e
}

func (e *ListBillingUsageMetric) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ComputeCredit":
		fallthrough
	case "StorageAvgByte":
		*e = ListBillingUsageMetric(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListBillingUsageMetric: %v", v)
	}
}

type ListBillingUsageRequest struct {
	// The aggregate type used to group usage which includes hour, day and month. default is hour
	AggregateBy *ListBillingUsageAggregateBy `queryParam:"style=form,explode=true,name=aggregateBy"`
	// The end time for the usage interval valid UTC ISO8601 format e.g. 2023-07-30T18:30:00Z
	EndTime string `queryParam:"style=form,explode=true,name=endTime"`
	// Metrics include ComputeCredit, StorageAvgByte. default is all
	Metric *ListBillingUsageMetric `queryParam:"style=form,explode=true,name=metric"`
	// The start time for the usage interval in valid UTC ISO8601 format e.g. 2023-07-30T18:30:00Z
	StartTime string `queryParam:"style=form,explode=true,name=startTime"`
}

type ListBillingUsageResponse struct {
	// OK
	BillingUsage *shared.BillingUsage
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
