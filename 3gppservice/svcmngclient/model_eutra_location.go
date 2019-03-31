/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package svcmngclient
import (
	"time"
)

type EutraLocation struct {
	Tai Tai `json:"tai"`
	Ecgi Ecgi `json:"ecgi"`
	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation string `json:"geographicalInformation,omitempty"`
	GeodeticInformation string `json:"geodeticInformation,omitempty"`
	GlobalNgenbId GlobalRanNodeId `json:"globalNgenbId,omitempty"`
}