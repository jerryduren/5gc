/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package svcmngserver

type SmfInfo struct {

	SNssaiSmfInfoList []SnssaiSmfInfoItem `json:"sNssaiSmfInfoList"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	PgwFqdn string `json:"pgwFqdn,omitempty"`

	AccessType []AccessType `json:"accessType,omitempty"`
}