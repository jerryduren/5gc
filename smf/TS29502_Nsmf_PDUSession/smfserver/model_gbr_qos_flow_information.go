/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smfserver

type GbrQosFlowInformation struct {

	MaxFbrDl string `json:"maxFbrDl"`

	MaxFbrUl string `json:"maxFbrUl"`

	GuaFbrDl string `json:"guaFbrDl"`

	GuaFbrUl string `json:"guaFbrUl"`

	NotifControl *NotificationControl `json:"notifControl,omitempty"`

	MaxPacketLossRateDl int32 `json:"maxPacketLossRateDl,omitempty"`

	MaxPacketLossRateUl int32 `json:"maxPacketLossRateUl,omitempty"`
}