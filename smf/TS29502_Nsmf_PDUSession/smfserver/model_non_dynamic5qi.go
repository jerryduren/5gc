/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smfserver

type NonDynamic5qi struct {

	PriorityLevel int32 `json:"priorityLevel,omitempty"`

	AverWindow string `json:"averWindow,omitempty"`

	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty"`
}
