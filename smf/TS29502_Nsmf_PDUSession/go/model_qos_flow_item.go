/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package go

type QosFlowItem struct {

	Qfi int32 `json:"qfi"`

	Cause *Cause `json:"cause,omitempty"`
}