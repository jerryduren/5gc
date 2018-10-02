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

type SmContextUpdateData struct {
	Pei string `json:"pei,omitempty" xml:"pei"`
	Gpsi string `json:"gpsi,omitempty" xml:"gpsi"`
	ServingNfId string `json:"servingNfId,omitempty" xml:"servingNfId"`
	Guami Guami `json:"guami,omitempty" xml:"guami"`
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty" xml:"backupAmfInfo"`
	AnType AccessType `json:"anType,omitempty" xml:"anType"`
	PresenceInLadn bool `json:"presenceInLadn,omitempty" xml:"presenceInLadn"`
	UeLocation UserLocation `json:"ueLocation,omitempty" xml:"ueLocation"`
	UeTimeZone string `json:"ueTimeZone,omitempty" xml:"ueTimeZone"`
	AddUeLocation UserLocation `json:"addUeLocation,omitempty" xml:"addUeLocation"`
	AddUeLocTime time.Time `json:"addUeLocTime,omitempty" xml:"addUeLocTime"`
	UpCnxState UpCnxState `json:"upCnxState,omitempty" xml:"upCnxState"`
	HoState HoState `json:"hoState,omitempty" xml:"hoState"`
	ToBeSwitched bool `json:"toBeSwitched,omitempty" xml:"toBeSwitched"`
	N1SmMsg RefToBinaryData `json:"n1SmMsg,omitempty" xml:"n1SmMsg"`
	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty" xml:"n2SmInfo"`
	N2SmInfoType int32 `json:"n2SmInfoType,omitempty" xml:"n2SmInfoType"`
	TargetServingNfId string `json:"targetServingNfId,omitempty" xml:"targetServingNfId"`
	SmContextStatusUri string `json:"smContextStatusUri,omitempty" xml:"smContextStatusUri"`
	DataForwarding bool `json:"dataForwarding,omitempty" xml:"dataForwarding"`
	EpsBearerSetup []string `json:"epsBearerSetup,omitempty" xml:"epsBearerSetup"`
	RevokeEbiList []int32 `json:"revokeEbiList,omitempty" xml:"revokeEbiList"`
	Release bool `json:"release,omitempty" xml:"release"`
	Cause Cause `json:"cause,omitempty" xml:"cause"`
}