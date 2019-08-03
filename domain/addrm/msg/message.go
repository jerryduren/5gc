package msg

import (
	"github.com/jerryduren/5gc/domain/cmddata"
	"github.com/jerryduren/5gc/domain/cmdmsg"
)


type GetAddrRequest struct {
	cmdmsg.MessageHead
	RequestAddrDomain cmddata.AddressDomain					`json:"request_addr_domain"`
	Dnn string																						`json:"dnn"`
	Location cmddata.UserLocation												`json:"location"`
	RequestedAddrType cmddata.AddressType						`json:"requested_addr_type"`
}

type GetAddrResponse struct {
	cmdmsg.MessageHead
	AllocatedAddrType cmddata.AddressType						`json:"allocated_addr_type"`
	AllocatedAddr	string																	`json:"allocated_addr"`
}

type ReleaseAddrRequest struct {

}

type ReleaseAddrResponse struct {

}