/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
import (
	"time"
)

type SmContextReleaseData struct {
	Cause Cause `json:"cause,omitempty" xml:"cause"`
	UeLocation UserLocation `json:"ueLocation,omitempty" xml:"ueLocation"`
	UeTimeZone string `json:"ueTimeZone,omitempty" xml:"ueTimeZone"`
	AddUeLocation UserLocation `json:"addUeLocation,omitempty" xml:"addUeLocation"`
	AddUeLocTime time.Time `json:"addUeLocTime,omitempty" xml:"addUeLocTime"`
	VsmfReleaseOnly bool `json:"vsmfReleaseOnly,omitempty" xml:"vsmfReleaseOnly"`
}
