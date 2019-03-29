/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ts29502nsmfpduclient

type PduSessionCreateData struct {
	Supi map[string]interface{} `json:"supi,omitempty"`
	UnauthenticatedSupi bool `json:"unauthenticatedSupi,omitempty"`
	Pei map[string]interface{} `json:"pei,omitempty"`
	PduSessionId map[string]interface{} `json:"pduSessionId,omitempty"`
	Dnn map[string]interface{} `json:"dnn"`
	SNssai map[string]interface{} `json:"sNssai,omitempty"`
	VsmfId map[string]interface{} `json:"vsmfId"`
	ServingNetwork map[string]interface{} `json:"servingNetwork"`
	RequestType RequestType `json:"requestType,omitempty"`
	EpsBearerId []int32 `json:"epsBearerId,omitempty"`
	PgwS8cFteid map[string]interface{} `json:"pgwS8cFteid,omitempty"`
	VsmfPduSessionUri map[string]interface{} `json:"vsmfPduSessionUri"`
	VcnTunnelInfo TunnelInfo `json:"vcnTunnelInfo"`
	AnType map[string]interface{} `json:"anType"`
	RatType map[string]interface{} `json:"ratType,omitempty"`
	UeLocation map[string]interface{} `json:"ueLocation,omitempty"`
	UeTimeZone map[string]interface{} `json:"ueTimeZone,omitempty"`
	AddUeLocation map[string]interface{} `json:"addUeLocation,omitempty"`
	Gpsi map[string]interface{} `json:"gpsi,omitempty"`
	N1SmInfoFromUe map[string]interface{} `json:"n1SmInfoFromUe,omitempty"`
	UnknownN1SmInfo map[string]interface{} `json:"unknownN1SmInfo,omitempty"`
	SupportedFeatures map[string]interface{} `json:"supportedFeatures,omitempty"`
	HPcfId map[string]interface{} `json:"hPcfId,omitempty"`
	HoPreparationIndication bool `json:"hoPreparationIndication,omitempty"`
	SelMode DnnSelectionMode `json:"selMode,omitempty"`
	AlwaysOnRequested bool `json:"alwaysOnRequested,omitempty"`
	UdmGroupId map[string]interface{} `json:"udmGroupId,omitempty"`
	RoutingIndicator string `json:"routingIndicator,omitempty"`
	EpsInterworkingInd EpsInterworkingIndication `json:"epsInterworkingInd,omitempty"`
	VSmfServiceInstanceId string `json:"vSmfServiceInstanceId,omitempty"`
	RecoveryTime map[string]interface{} `json:"recoveryTime,omitempty"`
	RoamingChargingProfile map[string]interface{} `json:"roamingChargingProfile,omitempty"`
	ChargingId string `json:"chargingId,omitempty"`
	OldPduSessionId map[string]interface{} `json:"oldPduSessionId,omitempty"`
}
