/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package go

type QosFlowProfile struct {

	Var5qi int32 `json:"5qi,omitempty"`

	NonDynamic5qi *NonDynamic5qi `json:"nonDynamic5qi,omitempty"`

	Dynamic5qi *Dynamic5qi `json:"dynamic5qi,omitempty"`

	Arp *Arp `json:"arp,omitempty"`

	GbrQosFlowInfo *GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`

	Rqa *ReflectiveQosAttribute `json:"rqa,omitempty"`
}
