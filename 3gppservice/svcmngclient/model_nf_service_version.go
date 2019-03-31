/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package svcmngclient
import (
	"time"
)

type NfServiceVersion struct {
	ApiVersionInUri string `json:"apiVersionInUri"`
	ApiFullVersion string `json:"apiFullVersion"`
	Expiry time.Time `json:"expiry,omitempty"`
}