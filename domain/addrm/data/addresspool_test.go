package data

import (
	"fmt"
	"testing"
)

func TestNewDnnAddressPool(t *testing.T) {
	dnnAddressPool:=NewDnnAddressPool("cmnetAddressPool")
	dnnAddressPool.AddAddressPool("","192.168.0.0")
	dnnAddressPool.AddAddressPool("","172.10.0.0")
	
	dnnAddressPool.AddAddressPool("460001234567","192.168.0.0")
	dnnAddressPool.AddAddressPool("460001234567","172.10.0.0")
	
	dnnAddressPool.AddAddressPool("460007890123","175.10.0.0")
	dnnAddressPool.AddAddressPool("460007890123","176.10.0.0")
	
	
	fmt.Println(dnnAddressPool)
	// 如何程序检查是否正确？
	// TODO
	
	
}

func TestNewAddressPool(t *testing.T) {
	addressPool:=NewAddressPool("192.168.0.0")
	addressPool.AddIpv4AddressSegment(ipv4ToUInt32("192.168.10.0"),8)
	addressPool.AddIpv4AddressSegment(ipv4ToUInt32("192.168.20.0"),16)
	
	fmt.Println("1:",addressPool)
	
	allocatedAddr:=make([]uint32,0)
	for {
		addr,err:=addressPool.ApplyIpv4Address()
		
		if err!=nil{
			fmt.Println("2: ", addressPool)
			t.Error(err)
			break
		}else {
			allocatedAddr=append(allocatedAddr,addr)
		}
	}
	
	for _,addr:=range allocatedAddr{
		_=addressPool.ReleaseIpv4Address(addr)
	}
	
	fmt.Println("3: ", addressPool)
	
	addr,err:=addressPool.ApplyIpv4Address()
	fmt.Println("Applied Addr: ",addr/*ipv4ToString(addr)*/,err)		// 因为刚释放的地址还没有间隔5s以上，申请失败
}