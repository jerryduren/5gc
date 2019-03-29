/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ts29502nsmfpduserver

type HsmfUpdateData struct {

	RequestIndication *RequestIndication `json:"requestIndication"`

	Pei *map[string]interface{} `json:"pei,omitempty"`

	VcnTunnelInfo *TunnelInfo `json:"vcnTunnelInfo,omitempty"`

	ServingNetwork *map[string]interface{} `json:"servingNetwork,omitempty"`

	AnType *map[string]interface{} `json:"anType,omitempty"`

	RatType *map[string]interface{} `json:"ratType,omitempty"`

	UeLocation *map[string]interface{} `json:"ueLocation,omitempty"`

	UeTimeZone *map[string]interface{} `json:"ueTimeZone,omitempty"`

	AddUeLocation *map[string]interface{} `json:"addUeLocation,omitempty"`

	PauseCharging bool `json:"pauseCharging,omitempty"`

	Pti int32 `json:"pti,omitempty"`

	N1SmInfoFromUe *map[string]interface{} `json:"n1SmInfoFromUe,omitempty"`

	UnknownN1SmInfo *map[string]interface{} `json:"unknownN1SmInfo,omitempty"`

	QosFlowsRelNotifyList []QosFlowItem `json:"qosFlowsRelNotifyList,omitempty"`

	QosFlowsNotifyList []QosFlowNotifyItem `json:"qosFlowsNotifyList,omitempty"`

	NotifyList []PduSessionNotifyItem `json:"NotifyList,omitempty"`

	EpsBearerId []int32 `json:"epsBearerId,omitempty"`

	HoPreparationIndication bool `json:"hoPreparationIndication,omitempty"`

	RevokeEbiList []int32 `json:"revokeEbiList,omitempty"`

	Cause *Cause `json:"cause,omitempty"`

	NgApCause *map[string]interface{} `json:"ngApCause,omitempty"`

	Var5gMmCauseValue *map[string]interface{} `json:"5gMmCauseValue,omitempty"`

	AlwaysOnRequested bool `json:"alwaysOnRequested,omitempty"`

	EpsInterworkingInd *EpsInterworkingIndication `json:"epsInterworkingInd,omitempty"`

	SecondaryRatUsageReport []map[string]interface{} `json:"secondaryRatUsageReport,omitempty"`
}
