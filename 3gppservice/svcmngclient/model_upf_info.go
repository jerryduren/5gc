/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package svcmngclient

type UpfInfo struct {
	SNssaiUpfInfoList []SnssaiUpfInfoItem `json:"sNssaiUpfInfoList"`
	SmfServingArea []string `json:"smfServingArea,omitempty"`
	InterfaceUpfInfoList []InterfaceUpfInfoItem `json:"interfaceUpfInfoList,omitempty"`
	IwkEpsInd bool `json:"iwkEpsInd,omitempty"`
}
