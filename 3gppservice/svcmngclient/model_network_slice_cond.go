/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package svcmngclient

type NetworkSliceCond struct {
	SnssaiList []Snssai `json:"snssaiList"`
	NsiList []string `json:"nsiList,omitempty"`
}
