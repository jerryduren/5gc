/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smfserver

type TunnelInfo struct {

	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addr string `json:"ipv6Addr,omitempty"`

	GtpTeid string `json:"gtpTeid"`
}
