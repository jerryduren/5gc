/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package svcmngserver

import (
	"time"
)

type NrLocationRm struct {

	Tai *Tai `json:"tai"`

	Ncgi *Ncgi `json:"ncgi"`

	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`

	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`

	GeographicalInformation string `json:"geographicalInformation,omitempty"`

	GeodeticInformation string `json:"geodeticInformation,omitempty"`

	GlobalGnbId *GlobalRanNodeId `json:"globalGnbId,omitempty"`
}
