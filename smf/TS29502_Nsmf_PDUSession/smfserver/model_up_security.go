/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smfserver

type UpSecurity struct {

	UpIntegr *UpIntegrity `json:"upIntegr"`

	UpConfid *UpConfidentiality `json:"upConfid"`
}
