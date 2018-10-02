/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
import (
	"time"
)

type SmContextCreateData struct {
	Supi string `json:"supi,omitempty" xml:"supi"`
	UnauthenticatedSupi bool `json:"unauthenticatedSupi,omitempty" xml:"unauthenticatedSupi"`
	Pei string `json:"pei,omitempty" xml:"pei"`
	Gpsi string `json:"gpsi,omitempty" xml:"gpsi"`
	PduSessionId int32 `json:"pduSessionId,omitempty" xml:"pduSessionId"`
	Dnn string `json:"dnn,omitempty" xml:"dnn"`
	SNssai Snssai `json:"sNssai,omitempty" xml:"sNssai"`
	HplmnSnssai Snssai `json:"hplmnSnssai,omitempty" xml:"hplmnSnssai"`
	ServingNfId string `json:"servingNfId" xml:"servingNfId"`
	Guami Guami `json:"guami,omitempty" xml:"guami"`
	RequestType RequestType `json:"requestType,omitempty" xml:"requestType"`
	N1SmMsg RefToBinaryData `json:"n1SmMsg,omitempty" xml:"n1SmMsg"`
	AnType AccessType `json:"anType" xml:"anType"`
	PresenceInLadn bool `json:"presenceInLadn,omitempty" xml:"presenceInLadn"`
	UeLocation UserLocation `json:"ueLocation,omitempty" xml:"ueLocation"`
	UeTimeZone string `json:"ueTimeZone,omitempty" xml:"ueTimeZone"`
	AddUeLocation UserLocation `json:"addUeLocation,omitempty" xml:"addUeLocation"`
	AddUeLocTime time.Time `json:"addUeLocTime,omitempty" xml:"addUeLocTime"`
	SmContextStatusUri string `json:"smContextStatusUri" xml:"smContextStatusUri"`
	HSmfUri string `json:"hSmfUri,omitempty" xml:"hSmfUri"`
	AdditionalHsmfUri []string `json:"additionalHsmfUri,omitempty" xml:"additionalHsmfUri"`
	OldPduSessionId int32 `json:"oldPduSessionId,omitempty" xml:"oldPduSessionId"`
	PduSessionsActivateList []int32 `json:"pduSessionsActivateList,omitempty" xml:"pduSessionsActivateList"`
	UeEpsPdnConnection string `json:"ueEpsPdnConnection,omitempty" xml:"ueEpsPdnConnection"`
	HoState HoState `json:"hoState,omitempty" xml:"hoState"`
	PcfId string `json:"pcfId,omitempty" xml:"pcfId"`
	SupportedFeatures string `json:"supportedFeatures,omitempty" xml:"supportedFeatures"`
	SelMode DnnSelectionMode `json:"selMode,omitempty" xml:"selMode"`
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty" xml:"backupAmfInfo"`
}
