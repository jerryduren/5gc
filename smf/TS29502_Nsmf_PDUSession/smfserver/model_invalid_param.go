/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smfserver

type InvalidParam struct {

	Param string `json:"param"`

	Reason string `json:"reason,omitempty"`
}