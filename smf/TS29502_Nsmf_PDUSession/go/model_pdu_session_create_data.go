/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package go

import (
	"time"
)

type PduSessionCreateData struct {

	Supi string `json:"supi,omitempty"`

	UnauthenticatedSupi bool `json:"unauthenticatedSupi,omitempty"`

	Pei string `json:"pei,omitempty"`

	PduSessionId int32 `json:"pduSessionId,omitempty"`

	Dnn string `json:"dnn"`

	SNssai *Snssai `json:"sNssai,omitempty"`

	VsmfId string `json:"vsmfId"`

	RequestType *RequestType `json:"requestType,omitempty"`

	EpsBearerId []int32 `json:"epsBearerId,omitempty"`

	PgwS8cFteid string `json:"pgwS8cFteid,omitempty"`

	VsmfPduSessionUri string `json:"vsmfPduSessionUri"`

	VcnTunnelInfo *TunnelInfo `json:"vcnTunnelInfo"`

	AnType *AccessType `json:"anType"`

	UeLocation *UserLocation `json:"ueLocation,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	AddUeLocation *UserLocation `json:"addUeLocation,omitempty"`

	AddUeLocTime time.Time `json:"addUeLocTime,omitempty"`

	Gpsi string `json:"gpsi,omitempty"`

	N1SmInfoFromUe *RefToBinaryData `json:"n1SmInfoFromUe,omitempty"`

	UnknownN1SmInfo *RefToBinaryData `json:"unknownN1SmInfo,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	HPcfId string `json:"hPcfId,omitempty"`

	HoPreparationIndication bool `json:"hoPreparationIndication,omitempty"`

	SelMode *DnnSelectionMode `json:"selMode,omitempty"`
}
