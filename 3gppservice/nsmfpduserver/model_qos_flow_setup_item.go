/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nsmfpduserver

type QosFlowSetupItem struct {

	Qfi *map[string]interface{} `json:"qfi"`

	QosRules *map[string]interface{} `json:"qosRules"`

	Ebi int32 `json:"ebi,omitempty"`

	QosFlowDescription *map[string]interface{} `json:"qosFlowDescription,omitempty"`

	QosFlowProfile *QosFlowProfile `json:"qosFlowProfile,omitempty"`
}