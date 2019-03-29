/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ts29502nsmfpduserver

type SmContextUpdateError struct {

	Error *map[string]interface{} `json:"error"`

	N1SmMsg *map[string]interface{} `json:"n1SmMsg,omitempty"`

	N2SmInfo *map[string]interface{} `json:"n2SmInfo,omitempty"`

	N2SmInfoType *N2SmInfoType `json:"n2SmInfoType,omitempty"`

	UpCnxState *UpCnxState `json:"upCnxState,omitempty"`

	RecoveryTime *map[string]interface{} `json:"recoveryTime,omitempty"`
}
