package msg

import (
	"github.com/jerryduren/5gc/3gppservice/commondata"
	"github.com/jerryduren/5gc/domain/cmdmsg"
)


type GetAddrRequest struct {
	cmdmsg.MessageHead
	RequestAddrDomain commondata.AddressDomain `json:"request_addr_domain"`
	Dnn               string                   `json:"dnn"`
	Location          commondata.UserLocation  `json:"location"`
	RequestedAddrType commondata.AddressType   `json:"requested_addr_type"`
}

type GetAddrResponse struct {
	cmdmsg.MessageHead
	AllocatedAddrType commondata.AddressType `json:"allocated_addr_type"`
	AllocatedAddr     string                 `json:"allocated_addr"`
}

type ReleaseAddrRequest struct {

}

type ReleaseAddrResponse struct {

}