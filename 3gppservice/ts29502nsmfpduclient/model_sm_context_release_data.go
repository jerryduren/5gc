/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ts29502nsmfpduclient

type SmContextReleaseData struct {
	Cause Cause `json:"cause,omitempty"`
	NgApCause map[string]interface{} `json:"ngApCause,omitempty"`
	Var5gMmCauseValue map[string]interface{} `json:"5gMmCauseValue,omitempty"`
	UeLocation map[string]interface{} `json:"ueLocation,omitempty"`
	UeTimeZone map[string]interface{} `json:"ueTimeZone,omitempty"`
	AddUeLocation map[string]interface{} `json:"addUeLocation,omitempty"`
	VsmfReleaseOnly bool `json:"vsmfReleaseOnly,omitempty"`
}
