/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smfserver

type PduSessionCreatedData struct {

	PduSessionType *PduSessionType `json:"pduSessionType"`

	SscMode string `json:"sscMode"`

	HcnTunnelInfo *TunnelInfo `json:"hcnTunnelInfo"`

	SessionAmbr *Ambr `json:"sessionAmbr"`

	QosFlowsSetupList []QosFlowSetupItem `json:"qosFlowsSetupList"`

	PduSessionId int32 `json:"pduSessionId,omitempty"`

	SNssai *Snssai `json:"sNssai,omitempty"`

	EnablePauseCharging bool `json:"enablePauseCharging,omitempty"`

	UeIpv4Address string `json:"ueIpv4Address,omitempty"`

	UeIpv6Prefix string `json:"ueIpv6Prefix,omitempty"`

	N1SmInfoToUe *RefToBinaryData `json:"n1SmInfoToUe,omitempty"`

	EpsPdnCnxInfo *EpsPdnCnxInfo `json:"epsPdnCnxInfo,omitempty"`

	EpsBearerInfo []EpsBearerInfo `json:"epsBearerInfo,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	UpSecurity *UpSecurity `json:"upSecurity,omitempty"`
}
