/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ts29502nsmfpduclient

type QosFlowAddModifyRequestItem struct {
	Qfi map[string]interface{} `json:"qfi"`
	Ebi int32 `json:"ebi,omitempty"`
	QosRules map[string]interface{} `json:"qosRules,omitempty"`
	QosFlowDescription map[string]interface{} `json:"qosFlowDescription,omitempty"`
	QosFlowProfile QosFlowProfile `json:"qosFlowProfile,omitempty"`
}