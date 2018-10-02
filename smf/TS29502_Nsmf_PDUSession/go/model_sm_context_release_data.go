/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package go

import (
	"time"
)

type SmContextReleaseData struct {

	Cause *Cause `json:"cause,omitempty"`

	UeLocation *UserLocation `json:"ueLocation,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	AddUeLocation *UserLocation `json:"addUeLocation,omitempty"`

	AddUeLocTime time.Time `json:"addUeLocTime,omitempty"`

	VsmfReleaseOnly bool `json:"vsmfReleaseOnly,omitempty"`
}
