/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ts29502nsmfpduclient

type QosFlowProfile struct {
	Var5qi map[string]interface{} `json:"5qi"`
	NonDynamic5Qi map[string]interface{} `json:"nonDynamic5Qi,omitempty"`
	Dynamic5Qi map[string]interface{} `json:"dynamic5Qi,omitempty"`
	Arp map[string]interface{} `json:"arp,omitempty"`
	GbrQosFlowInfo GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
	Rqa map[string]interface{} `json:"rqa,omitempty"`
	AdditionalQosFlowInfo map[string]interface{} `json:"additionalQosFlowInfo,omitempty"`
}
