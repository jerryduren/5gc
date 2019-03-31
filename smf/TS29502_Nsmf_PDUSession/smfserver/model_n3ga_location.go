/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smfserver

type N3gaLocation struct {

	N3gppTai *Tai `json:"n3gppTai"`

	N3IwfId string `json:"n3IwfId"`

	UeIpv4Addr string `json:"ueIpv4Addr,omitempty"`

	UeIpv6Addr string `json:"ueIpv6Addr,omitempty"`

	PortNumber int32 `json:"portNumber,omitempty"`
}