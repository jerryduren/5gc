package nrf


type NFInstance struct {
	NFProfile
	// TODO
}

type NFInstances struct {
	nfInstances map[string]NFInstance
}

type CreateNFInstanceData struct {

}

type NFType string
const (
	NFTYPE_NRF NFType = "NRF"
	NFTYPE_UDM NFType = "UDM"
	NFTYPE_AMF NFType = "AMF"
	NFTYPE_SMF NFType = "SMF"
	NFTYPE_AUSF NFType = "AUSF"
	NFTYPE_NEF NFType = "NEF"
	NFTYPE_PCF NFType = "PCF"
	NFTYPE_SMSF NFType = "SMCF"
	NFTYPE_NSSF NFType = "NSSF"
	NFTYPE_UDR NFType = "UDR"
	NFTYPE_LMF NFType = "LMF"
	NFTYPE_GMLC NFType = "GMLC"
	NFTYPE_5G_EIR NFType = "5G_EIR"
	NFTYPE_SEPP NFType = "SEPP"
	NFTYPE_NFT NFType = "NFT"
	NFTYPE_UPF NFType = "UPF"
	NFTYPE_N3IWF NFType = "N3IWF"
	NFTYPE_AF NFType = "AF"
	NFTYPE_UDSF NFType = "UDCF"
	NFTYPE_BSF NFType = "BSF"
	NFTYPE_CHF NFType = "CHF"
	NFTYPE_NWDAF NFType = "NWDAF"
)

type NFStatus string
const (
	NFSTATUS_REGISTERED NFStatus = "REGISTERED"
	NFSTATUS_SUSPENDED NFStatus = "SUSPENDED"
	NFSTATUS_UNDISCOVERABLE NFStatus = "UNDISCOVERABLE"
)

type PlmnId struct {
	Mcc string		`json:"mcc"`			// ^\d{3}$
	Mnc string 		`json:"mnc"`			// ^\d{2,3}$
}
func (plmn PlmnId)Strings()string{
	return plmn.Mcc+plmn.Mnc
}

// Single Network Slice Select Assistant Information
type Snssai struct {
	Sst int32		`json:"sst"`			// 0~255, slice service type
	Sd string		`json:"sd"`				// ^[A-Fa-f0-9]{6}$   ,service difference
}

type PerPlmnSnssai struct{
	PlmnId PlmnId				`json:"plmnId"`
	SnssaiList []Snssai			`json:"snssaiList"`
}

type Nsi string					// Network Slice Instance

type UdrInfo struct {
	// TODO
}
type NfService	struct{

//	TODO
}

type DefaultNotificationSubscription struct {
//	TODO
}

type NFProfile struct {
	// required
	NfInstanceId		string				`json:"nfInstanceId"`		// uuid
	NfType				NFType				`json:"nfType"`				// enum
	NfStatus			NFStatus			`json:"nfStatus"`			// enum
	
	// options
	HeartBeatTimer		int32				`json:"heartBeatTimer,omitempty"`		// heartbeat timer, unit: second
	PlmnList			[]PlmnId			`json:"plmnList,omitempty"`				// Plmn list
	Snssais				[]Snssai			`json:"snssais,omitempty"`				//
	PerPlmnSnssaiList 	[]PerPlmnSnssai		`json:"perPlmnSnssaiList,omitempty"`	//
	NsiList				[]Nsi				`json:"nsiList,omitempty"`
	Fqdn				string				`json:"fqdn,omitempty"`
	InterPlmnFqdn		string				`json:"interPlmnFqdn,omitempty"`
	Ipv4Addresses		[]string			`json:"ipv4Addresses,omitempty"`		// ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$  example: 192.168.10.1
	Ipv6Addresses		[]string			`json:"ipv6Addresses,omitempty"`		// ^((([^:]+:){7}([^:]+))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?))$ example: 2001:db8:85a3::8a2e:370:7334
	AllowedPlmns		[]PlmnId			`json:"allowedPlmns,omitempty"`
	AllowedNfTypes		[]NFType			`json:"allowedNfTypes,omitempty"`
	AllowedNfDomains	[]string			`json:"allowedNfDomains,omitempty"`				// what is domain of NF?
	AllowedNssais		[]Snssai			`json:"allowedNssais,omitempty"`
	Priority			int32				`json:"priority,omitempty"`						// 0~65535
	Capacity			int32				`json:"capacity,omitempty"`						// 0~65535
	Load				int32				`json:"load,omitempty"`							// 0~100
	Locality			string				`json:"locality,omitempty"`
	UdrInfo				UdrInfo				`json:"udrInfo,omitempty"`
	// TODO
	CustomInfo			string				`json:"customInfo,omitempty"`
	RecoveryTime		string				`json:"recoveryTime,omitempty"`					// $date-time
	NfServicePersistence bool				`json:"nfServicePersistence,omitempty"`
	NfServices			[]NfService			`json:"nfServices,omitempty"`
	NfProfileChangesSupportInd	bool		`json:"nfProfileChangesSupportInd,omitempty"`
	NfProfileChangesInd	bool				`json:"nfProfileChangesInd,omitempty"`
	DefaultNotificationSubscriptions	[]DefaultNotificationSubscription	`json:"defaultNotificationSubscriptions,omitempty"`
}

func (nfIns *NFInstances)RegisterNFInstance(nfInstanceID string, createNFInstanceData CreateNFInstanceData) error{
	return  nil
}

func (nfIns *NFInstances)GetNFInstance(nfInstanceID string)NFInstance{
	return nfIns.nfInstances[nfInstanceID]
}