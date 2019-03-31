/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ts29502nsmfpduclient

type SmContextUpdateError struct {
	Error map[string]interface{} `json:"error"`
	N1SmMsg map[string]interface{} `json:"n1SmMsg,omitempty"`
	N2SmInfo map[string]interface{} `json:"n2SmInfo,omitempty"`
	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`
	UpCnxState UpCnxState `json:"upCnxState,omitempty"`
	RecoveryTime map[string]interface{} `json:"recoveryTime,omitempty"`
}