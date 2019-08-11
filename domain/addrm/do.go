package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	
	"github.com/jerryduren/5gc/domain/addrm/context"
	"github.com/jerryduren/5gc/domain/addrm/data"
)

func main(){
	if DemoApplyIpv4AddressService()!=nil{
		fmt.Println("FAILURE")
	}
}
// Example
func DemoApplyIpv4AddressService() error{
	admService:=context.NewAdmService()
	
	//  初始化两个仓库
	dnnAddrPool:=data.NewDnnAddressPool("cmnet")
	addrPool:=data.NewAddressPool("192.168.0.0")
	
	dnnAddrPool.AddAddressPool("", addrPool.Name)
	dnnAddrPool.AddAddressPool("46005123456789",addrPool.Name)
	
	addrPool.AddIpv4AddressSegment(ipv4ToUInt32("192.168.10.1"),8)
	
	_=admService.DnnAddrPoolRepository.Add(dnnAddrPool)
	_=admService.AddrPoolRepository.Add(addrPool)
	
	if addr,err:=admService.ApplyIpv4Address("cmnet","");err!=nil{
		return errors.New("Apply Address Failure")
	}else {
		fmt.Println("Apply Address:",ipv4ToString(addr))
	}
	
	if addr,err:=admService.ApplyIpv4Address("cmnet","46005123456789");err!=nil{
		return errors.New("Apply Address Failure")
	}else {
		fmt.Println("Apply Address:",ipv4ToString(addr))
	}
	
	return nil
}


// the format of ipv4Addr is same as 192.168.10.3
func ipv4ToUInt32(ipv4 string) uint32{
	ps:=strings.Split(ipv4,".")
	if len(ps)>=4{
		p1,_:=strconv.Atoi(ps[0])
		p2,_:=strconv.Atoi(ps[1])
		p3,_:=strconv.Atoi(ps[2])
		p4,_:=strconv.Atoi(ps[3])
		return (uint32(p1)<<24)|(uint32(p2)<<16)|(uint32(p3)<<8)|uint32(p4)
	}
	
	return 0
}

func ipv4ToString(addr uint32)string{
	return fmt.Sprintf("%d.%d.%d.%d",addr&0xff000000>>24,addr&0x00ff0000>>16,addr&0x0000ff00>>8,addr&0x000000ff)
}