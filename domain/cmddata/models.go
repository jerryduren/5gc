package cmddata

import "time"

type AccessType string
// List of AccessType
const (
	AT_3GPP AccessType = "3GPP_ACCESS"
	AT_NON_3GPP AccessType = "NON_3GPP_ACCESS"
)

type AddressDomain string

type AddressType string
const(
	ADDR_TYPE_IPV4 AddressType = "IPv4"
	ADDR_TYPE_IPV6 AddressType = "IPv6"
)

type RatType string
// List of RatType
const (
	RAT_NR RatType = " NR"
	RAT_EUTRA RatType = "EUTRA"
	RAT_WLAN RatType = "WLAN"
	RAT_VIRTUAL RatType = "VIRTUAL"
)

type PduSessionType string
// List of PduSessionType
const (
	PDU_TYPE_IPV4 PduSessionType = "IPV4"
	PDU_TYPE_IPV4_IPV6 PduSessionType = "IPV6"
	PDU_TYPE_IPV4_IPV4V6 PduSessionType = "IPV4V6"
	PDU_TYPE_IPV4_UNSTRUCTURED PduSessionType = "UNSTRUCTURED"
	PDU_TYPE_IPV4_ETHERNET PduSessionType = "ETHERNET"
)


type UpIntegrity string
// List of UpIntegrity
const (
	UP_INT_REQUIRED UpIntegrity = "REQUIRED"
	UP_INT_PREFERRED UpIntegrity ="PREFERRED"
	UP_INT_NOT_NEEDED UpIntegrity ="NOT_NEEDED"
)

type UpConfidentiality string
// List of UpConfidentiality
const (
	UP_CFD_REQUIRED UpConfidentiality = "REQUIRED"
	UP_CFD_PREFERRED UpConfidentiality ="PREFERRED"
	UP_CFD_NOT_NEEDED UpConfidentiality ="NOT_NEEDED"
)

type SscMode string
// List of SscMode
const (
	SSC_MODE_1 SscMode ="SSC_MODE_1"
	SSC_MODE_2 SscMode= "SSC_MODE_2"
	SSC_MODE_3 SscMode ="SSC_MODE_3"
)


type DnaiChangeType string
// List of DnaiChangeType
/*description: >
This string provides forward-compatibility with future
extensions to the enumeration but is not used to encode
content defined in the present version of this API.
description: >
Possible values are
- EARLY: Early notification of UP path reconfiguration.
- EARLY_LATE: Early and late notification of UP path reconfiguration. This value shall only be present in the subscription to the DNAI change event.
- LATE: Late notification of UP path reconfiguration.
*/
const (
	DNAI_CHN_TYOE_EARLY  DnaiChangeType ="EARLY "
	DNAI_CHN_TYOE_EARLY_LATE DnaiChangeType= "EARLY_LATE"
	DNAI_CHN_TYOE_LATE DnaiChangeType ="LATE"
)

type RestrictionType string
// List of RestrictionType
const (
	RST_TYPE_ALLOWED_AREAS RestrictionType ="ALLOWED_AREAS"
	RST_TYPE_NOT_ALLOWED_AREAS RestrictionType= "NOT_ALLOWED_AREAS"
)

type CoreNetworkType string
// List of CoreNetworkType
const (
	CN_TYPE_5GC CoreNetworkType ="5GC"
	CN_TYPE_EPC CoreNetworkType= "EPC"
)

type PresenceState string
// List of PresenceState
const (
	PRS_STATE_IN_AREA PresenceState ="IN_AREA"
	PRS_STATE_OUT_OF_AREA PresenceState ="OUT_OF_AREA"
	PRS_STATE_UNKNOWN PresenceState ="UNKNOWN"
	PRS_STATE_INACTIVE PresenceState = "INACTIVE"
)

type ChargingId uint32

type RatingGroup uint32

type ServiceId uint32

type Supi string	//pattern: '^(imsi-[0-9]{5,15}|nai-.+|.+)$'

type PreemptionCapability string
const (
	PRM_CAP_NOT_PREEMPT PreemptionCapability = "NOT_PREEMPT"
	PRM_CAP_MAY_PREEMPT PreemptionCapability = "MAY_PREEMPT"
)

type PreemptionVulnerability string
const (
	PRM_VUL_NOT_PREEMPTABLE PreemptionVulnerability = "NOT_PREEMPTABLE"
	PRM_VUL_PREEMPTABLE PreemptionVulnerability = "PREEMPTABLE"
)

type ReflectiveQoSAttribute string
const (
	RQOS_RQOS ReflectiveQoSAttribute = "RQOS"
	RQOS_NO_RQOS ReflectiveQoSAttribute ="NO_RQOS"
)

type NotificationControl string
const (
	NTF_CTL_REQUESTED NotificationControl = "REQUESTED"
	NTF_CTL_NOT_REQUESTED NotificationControl = "NOT_REQUESTED"
)

type QosResourceType string
const (
	QOS_TYPE_NON_GBR QosResourceType = "NON_GBR"
	QOS_TYPE_NON_CRITICAL_GBR QosResourceType = "NON_CRITICAL_GBR"
	QOS_TYPE_CRITICAL_GBR  QosResourceType = "CRITICAL_GBR"
)

type AdditionalQosFlowInfo string
const (
	ADDITION_QOS_INFO_MORE_LIKELY AdditionalQosFlowInfo = "MORE_LIKELY"
)

type Qfi uint32
const (
	MIN_QFI Qfi = 0
	MAX_QFI Qfi = 63
)

type FiveQi uint32
const (
	MIN_5QI FiveQi = 0
	MAX_5GI FiveQi = 255
)

type BitRate string		// pattern: '^\d+(\.\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$'

type ArpPriorityLevel uint32
const (
	MIN_ARP_PRIORITY_LEVEL ArpPriorityLevel = 1
	MAX_ARP_PRIORITY_LEVEL ArpPriorityLevel = 15
)

type FiveQiPriorityLevel uint32
const (
	MIN_5QI_PRIORITY_LEVEL FiveQiPriorityLevel = 1
	MAX_5QI_PRIORITY_LEVEL FiveQiPriorityLevel = 127
	DEFAULT_5QI_PRIORITY_LEVEL FiveQiPriorityLevel = 100
)

type PacketDelBudget uint32
const (
	MIX_PACKET_DEL_BUDGET PacketDelBudget = 1
)

type PacketErrRate string		// pattern: '^([0-9]E-[0-9])$'

type PacketLossRate uint32
const (
	MIN_PACKET_LOSS_RATE PacketLossRate = 0
	MAX_PACKET_LOSS_RATE PacketLossRate = 0
)

type AverWindow uint32
const (
	MIN_AVER_WINDOW AverWindow = 1
	MAX_AVER_WINDOW AverWindow = 4095
	DEFAULT_AVER_WINDOW AverWindow = 2000
)

type MaxDataBurstVol uint32
const (
	_MIN_MAX_DATA_BURST_VOL MaxDataBurstVol = 1
	_MAX_MAX_DATA_BURST_VOL MaxDataBurstVol = 4095
)

type Snssai struct {
	Sst int32 `json:"sst"`
	Sd string `json:"sd,omitempty"`
}

type PlmnId struct {
	Mcc string `json:"mcc"`
	Mnc string `json:"mnc"`
}

type Tai struct {
	PlmnId *PlmnId `json:"plmnId"`
	Tac string `json:"tac"`     // may be Tac
}

type Ecgi struct {
	PlmnId *PlmnId `json:"plmnId"`
	EutraCellId string `json:"eutraCellId"`
}

type Ncgi struct {
	PlmnId *PlmnId `json:"plmnId"`
	NrCellId string `json:"nrCellId"`
}

type UserLocation struct {
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty"`
	NrLocation *NrLocation `json:"nrLocation,omitempty"`
	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`
}

type EutraLocation struct {
	Tai *Tai `json:"tai"`
	Ecgi *Ecgi `json:"ecgi"`
	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation string `json:"geographicalInformation,omitempty"`
	GeodeticInformation string `json:"geodeticInformation,omitempty"`
	GlobalNgenbId *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
}

type NrLocation struct {
	Tai *Tai `json:"tai"`
	Ncgi *Ncgi `json:"ncgi"`
	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation string `json:"geographicalInformation,omitempty"`
	GeodeticInformation string `json:"geodeticInformation,omitempty"`
	GlobalGnbId *GlobalRanNodeId `json:"globalGnbId,omitempty"`
}

type N3gaLocation struct {
	N3gppTai *Tai `json:"n3gppTai,omitempty"`
	N3IwfId string `json:"n3IwfId,omitempty"`
	UeIpv4Addr string `json:"ueIpv4Addr,omitempty"`
	UeIpv6Addr string `json:"ueIpv6Addr,omitempty"`
	PortNumber int32 `json:"portNumber,omitempty"`
}

type UpSecurity struct {
	UpIntegr *UpIntegrity `json:"upIntegr"`
	UpConfid *UpConfidentiality `json:"upConfid"`
}

type NgApCause struct {
	Group int32 `json:"group"`
	Value int32 `json:"value"`
}

type AmfName string

type AmfId string		// pattern: '^[A-Fa-f0-9]{6}$'

type Guami struct {
	PlmnId PlmnId
	AmfId AmfId
}

type BackupAmfInfo struct {
	BackupAmf AmfName `json:"backupAmf"`
	GuamiList []Guami `json:"guamiList,omitempty"`
}

type RefToBinaryData struct {
	ContentId string `json:"contentId"`
}

type Dnai string

type RouteInformation struct {
	Ipv4Addr string `json:"ipv4Addr,omitempty"`
	Ipv6Addr string `json:"ipv6Addr,omitempty"`
	PortNumber int32 `json:"portNumber"`
}

type RouteToLocation struct {
	Dnai Dnai `json:"dnai"`
	RouteInfo *RouteInformation `json:"routeInfo,omitempty"`
	RouteProfId string `json:"routeProfId,omitempty"`
}

type SubscribedDefaultQos struct {
	Var5qi  FiveQi `json:"5qi"`
	Arp *Arp `json:"arp"`
	PriorityLevel FiveQiPriorityLevel `json:"priorityLevel,omitempty"`
}

type Tac string		// pattern: '(^[A-Fa-f0-9]{4}$)|(^[A-Fa-f0-9]{6}$)'

type AreaCode string		// union

type Area struct {
	Tacs []Tac
	AreaCodes []AreaCode
}

// 1st condition: restrictionType and areas attributes shall be either both absent or both present
type ServiceAreaRestriction struct {
	RestrictionType RestrictionType
	Areas []Area
	MaxNumOfTAs uint32
}

type PresenceInfo struct {
	PraId string `json:"praId,omitempty"`
	PresenceState *PresenceState `json:"presenceState,omitempty"`
	TrackingAreaList []Tai `json:"trackingAreaList,omitempty"`
	EcgiList []Ecgi `json:"ecgiList,omitempty"`
	NcgiList []Ncgi `json:"ncgiList,omitempty"`
	GlobalRanNodeIdList []GlobalRanNodeId `json:"globalRanNodeIdList,omitempty"`
}

type N3IwfId string		// pattern: '^[A-Fa-f0-9]+$'

type NgeNbId string		// pattern: '^(MacroNGeNB-[A-Fa-f0-9]{5}|LMacroNGeNB-[A-Fa-f0-9]{6}|SMacroNGeNB-[A-Fa-f0-9]{5})$'

type GlobalRanNodeId struct {
	plmnId PlmnId
    n3IwfId N3IwfId
    gNbId GNbId
    ngeNbId NgeNbId
}

type GNbId struct {
	BitLength int32 `json:"bitLength"`
	GNBValue string `json:"gNBValue"`
}

type Arp struct {
	PriorityLevel int32 `json:"priorityLevel"`
	PreemptCap *PreemptionCapability `json:"preemptCap"`
	PreemptVuln *PreemptionVulnerability `json:"preemptVuln"`
}

type Ambr struct {
	Uplink BitRate `json:"uplink"`
	Downlink BitRate `json:"downlink"`
}

type Dynamic5Qi struct {
	ResourceType *QosResourceType `json:"resourceType"`
	PriorityLevel FiveQiPriorityLevel `json:"priorityLevel"`
	PacketDelayBudget PacketDelBudget `json:"packetDelayBudget"`
	PacketErrRate PacketErrRate `json:"packetErrRate"`
	AverWindow AverWindow `json:"averWindow,omitempty"`
	MaxDataBurstVol MaxDataBurstVol `json:"maxDataBurstVol,omitempty"`
}

type NonDynamic5Qi struct {
	PriorityLevel FiveQiPriorityLevel `json:"priorityLevel,omitempty"`
	AverWindow AverWindow `json:"averWindow,omitempty"`
	MaxDataBurstVol MaxDataBurstVol `json:"maxDataBurstVol,omitempty"`
}

type TraceData struct {
	TraceRef string `json:"traceRef"`
	TraceDepth *TraceDepth `json:"traceDepth"`
	NeTypeList string `json:"neTypeList"`
	EventList string `json:"eventList"`
	CollectionEntityIpv4Addr string `json:"collectionEntityIpv4Addr,omitempty"`
	CollectionEntityIpv6Addr string `json:"collectionEntityIpv6Addr,omitempty"`
	InterfaceList string `json:"interfaceList,omitempty"`
}

type TraceDepth string
const (
	TRC_DEPTH_MINIMUM TraceDepth = "MINIMUM"
	TRC_DEPTH_MEDIUM TraceDepth = "MEDIUM"
	TRC_DEPTH_MAXIMUM TraceDepth = "MAXIMUM"
	TRC_DEPTH_MINIMUM_WO_VENDOR_EXTENSION TraceDepth = "MINIMUM_WO_VENDOR_EXTENSION"
	TRC_DEPTH_MEDIUM_WO_VENDOR_EXTENSION TraceDepth = "MEDIUM_WO_VENDOR_EXTENSION"
	TRC_DEPTH_MAXIMUM_WO_VENDOR_EXTENSION TraceDepth = "MAXIMUM_WO_VENDOR_EXTENSION"
)


type OdbData struct {
	RoamingOdb *RoamingOdb `json:"roamingOdb,omitempty"`
	OdbPacketServices *OdbPacketServices `json:"odbPacketServices,omitempty"`
}

type OdbPacketServices string
const (
	ODB_ALL_PACKET_SERVICES OdbPacketServices = "ALL_PACKET_SERVICES"
	ODB_ROAMER_ACCESS_HPLMN_AP OdbPacketServices = "ROAMER_ACCESS_HPLMN_AP"
	ODB_ROAMER_ACCESS_VPLMN_AP OdbPacketServices = "ROAMER_ACCESS_VPLMN_AP"
)

type RoamingOdb string
const (
	RODB_OUTSIDE_HOME_PLMN RoamingOdb = "OUTSIDE_HOME_PLMN"
	RODB_OUTSIDE_HOME_PLMN_COUNTRY RoamingOdb = "OUTSIDE_HOME_PLMN_COUNTRY"
)

type SecondaryRatUsageReport struct {
	SecondaryRatType *RatType `json:"secondaryRatType"`
	QosFlowsUsageData []QosFlowUsageReport `json:"qosFlowsUsageData"`
}

type QosFlowUsageReport struct {
	Qfi Qfi `json:"qfi"`
	StartTimeStamp time.Time `json:"startTimeStamp"`
	EndTimeStamp time.Time `json:"endTimeStamp"`
	DownlinkVolume int64 `json:"downlinkVolume"`
	UplinkVolume int64 `json:"uplinkVolume"`
}