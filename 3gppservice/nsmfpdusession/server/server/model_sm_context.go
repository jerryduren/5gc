/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0.alpha-1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

type SmContext struct {

	Supi *map[string]interface{} `json:"supi,omitempty"`

	UnauthenticatedSupi bool `json:"unauthenticatedSupi,omitempty"`

	Pei *map[string]interface{} `json:"pei,omitempty"`

	Gpsi *map[string]interface{} `json:"gpsi,omitempty"`

	PduSessionId *map[string]interface{} `json:"pduSessionId"`

	Dnn *map[string]interface{} `json:"dnn"`

	SNssai *map[string]interface{} `json:"sNssai"`

	HplmnSnssai *map[string]interface{} `json:"hplmnSnssai,omitempty"`

	ServingNfId *map[string]interface{} `json:"servingNfId"`

	Guami *map[string]interface{} `json:"guami,omitempty"`

	ServiceName *map[string]interface{} `json:"serviceName,omitempty"`

	ServingNetwork *map[string]interface{} `json:"servingNetwork"`

	AnType *map[string]interface{} `json:"anType"`

	RatType *map[string]interface{} `json:"ratType,omitempty"`

	SmContextStatusUri *map[string]interface{} `json:"smContextStatusUri"`

	HSmfUri *map[string]interface{} `json:"hSmfUri,omitempty"`

	PcfId *map[string]interface{} `json:"pcfId,omitempty"`

	SelMode *DnnSelectionMode `json:"selMode,omitempty"`

	TraceData *map[string]interface{} `json:"traceData,omitempty"`

	UdmGroupId *map[string]interface{} `json:"udmGroupId,omitempty"`

	RoutingIndicator string `json:"routingIndicator,omitempty"`

	EpsInterworkingInd *EpsInterworkingIndication `json:"epsInterworkingInd,omitempty"`

	PduSessionType *map[string]interface{} `json:"pduSessionType"`

	SscMode string `json:"sscMode"`

	SessionAmbr *map[string]interface{} `json:"sessionAmbr"`

	QosFlowsSetupList []QosFlowSetupItem `json:"qosFlowsSetupList"`

	HSmfInstanceId *map[string]interface{} `json:"hSmfInstanceId,omitempty"`

	EnablePauseCharging bool `json:"enablePauseCharging,omitempty"`

	UeIpv4Address *map[string]interface{} `json:"ueIpv4Address,omitempty"`

	UeIpv6Prefix *map[string]interface{} `json:"ueIpv6Prefix,omitempty"`

	EpsPdnCnxInfo *EpsPdnCnxInfo `json:"epsPdnCnxInfo,omitempty"`

	EpsBearerInfo []EpsBearerInfo `json:"epsBearerInfo,omitempty"`

	MaxIntegrityProtectedDataRate *MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRate,omitempty"`

	AlwaysOnGranted bool `json:"alwaysOnGranted,omitempty"`

	UpSecurity *map[string]interface{} `json:"upSecurity,omitempty"`

	HSmfServiceInstanceId string `json:"hSmfServiceInstanceId,omitempty"`

	RecoveryTime *map[string]interface{} `json:"recoveryTime,omitempty"`
}