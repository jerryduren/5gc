/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package go

import (
	"os"
)

type InlineResponse4002 struct {

	JsonData *PduSessionCreateError `json:"jsonData,omitempty"`

	BinaryDataN1SmInfoToUe **os.File `json:"binaryDataN1SmInfoToUe,omitempty"`
}
