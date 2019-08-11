package data

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	
)

var errNotAvailableAddress  = errors.New("Not Available Address")
var errAddressReleasedNotExists = errors.New("Address Released Not Exists")

type AddressStatus struct {
	AllocateTime time.Time
	ReleaseTime time.Time
	IsAllocated bool
}

type Ipv4AddressSegment struct {
	StartAddress uint32
	Number uint32
	StatusList []AddressStatus									// 每个地址的状态，包括是否已经分配，分配的时间戳
}

func (as Ipv4AddressSegment)ApplyAddress()(uint32,error) {
	// 释放后5s(后续可配置)才能重新分配出去
	for index, value := range as.StatusList {
		if !value.IsAllocated {
			if time.Since(as.StatusList[index].ReleaseTime).Seconds() > float64(5){
				as.StatusList[index].IsAllocated = true
				as.StatusList[index].AllocateTime = time.Now()
				return as.StartAddress + uint32(index), nil
			}
		}
	}
	
	return 0,errNotAvailableAddress
}

func (as Ipv4AddressSegment)ReleaseAddress(addr uint32)error{
	if addr< as.StartAddress || addr>=as.StartAddress+as.Number{
		return  errAddressReleasedNotExists
	}
	
	as.StatusList[int(addr-as.StartAddress)].IsAllocated = false
	as.StatusList[int(addr-as.StartAddress)].ReleaseTime = time.Now()
	
	return nil
}

func (as *Ipv4AddressSegment)ModifyAddressSegment(startAddress uint32,number uint32)error{
	// TODO
	/*if startAddress > as.StartAddress && startAddress+number < as.StartAddress+as.Number{
		for idx,v:=range as.StatusList[0
	}*/
	
	return nil
}

type  Ipv4AddressSegments []Ipv4AddressSegment
func NewIpv4AddressSegments()Ipv4AddressSegments{
	return make([]Ipv4AddressSegment,0)
}

func (as Ipv4AddressSegments)Less(i,j int)bool{
	if i>=len(as) || j>=len(as)||i<0 || j <0{
		return false
	}
	
	return as[i].StartAddress < as[j].StartAddress
}

func (as Ipv4AddressSegments)Len()int{
	return len(as)
}

func (as Ipv4AddressSegments)Swap(i,j int){
	as[i],as[j] = as[j],as[i]
}

type Ipv6AddressSegment struct {
	StartAddress string
	Number uint32
	Available []AddressStatus
}

type Ipv6AddressSegments []Ipv6AddressSegment
func NewIpv6AddressSegments()Ipv6AddressSegments{
	return make([]Ipv6AddressSegment,0)
}

type AddressPool struct {
	Name string
	Ipv4SortedSegmentList Ipv4AddressSegments
	Ipv6SortedSegmentList Ipv6AddressSegments
}

// implement Entity interface
func (ap AddressPool)GetKey()interface{}{
	return ap.Name
}

func (ap AddressPool)Less(v1,v2 interface{})bool{
	return v1.(string)<v2.(string)
}

func (ap AddressPool)String()string{
	result:=fmt.Sprintf("Name = %s: \n",ap.Name)
	for _,v:=range ap.Ipv4SortedSegmentList{
		result = result+fmt.Sprintf("---- %s; Number: %d; Status: ", ipv4ToString(v.StartAddress),v.Number)
		for _,v:=range v.StatusList{
			if v.IsAllocated{
				result = result + "T"
			}else {
				result = result + "F"
			}
		}
		result = result + "\n"
	}
	
	return result
}
// implement Entity interface end

func (ap *AddressPool)AddIpv4AddressSegment(ipv4StartAddr uint32, number uint32){
	ipv4AddrSeg := Ipv4AddressSegment{StartAddress:ipv4StartAddr,Number:number,StatusList:make([]AddressStatus,number)}
	ap.Ipv4SortedSegmentList=append(ap.Ipv4SortedSegmentList,ipv4AddrSeg)
	sort.Sort(ap.Ipv4SortedSegmentList)
	// 考虑如何避免地址段重复
	// TODO
}

func (ap *AddressPool)DeleteIpv4AddressSegment(){
	// TODO
}

func (ap *AddressPool)ApplyIpv4Address()(uint32,error){
	if len(ap.Ipv4SortedSegmentList)==0{
		return 0,errNotAvailableAddress
	}
	
	for _,value:=range ap.Ipv4SortedSegmentList{
		if newAddr,err:=value.ApplyAddress();err==nil{
			return newAddr,err
		}
	}
	
	return 0,errNotAvailableAddress
}

func (ap *AddressPool)ReleaseIpv4Address(addr uint32)error{
	if len(ap.Ipv4SortedSegmentList)==0{
		return errAddressReleasedNotExists
	}
	
	for _,value:=range ap.Ipv4SortedSegmentList{
		if value.ReleaseAddress(addr)==nil{
			return nil
		}
	}
	
	return errAddressReleasedNotExists
}

func NewAddressPool(name string)*AddressPool{
	// 应该要检查，Name不要为空
	return &AddressPool{
		Name:				                 name,
		Ipv4SortedSegmentList: NewIpv4AddressSegments(),
		Ipv6SortedSegmentList: NewIpv6AddressSegments(),
	}
}

type LocationAddressPool struct {
	Tai string															// plmn id + tac: 46000234123456
	AddressPoolNameList map[string]bool					// 位置关联的地址池名称列表
}

func (lap *LocationAddressPool)AddAddressPool(name string){
	lap.AddressPoolNameList[name] = true
}

func (lap *LocationAddressPool)DeleteAddressPool(name string){
	// TODO
}

type DnnAddressPool struct {
	Dnn string
	LocationList map[string]LocationAddressPool						// key is LocationAddressPool.Tai
	AddressPoolNameList map[string]bool
}

// implement Entity interface BEGIN ...注意，Entity的实现不能是指针方法，类型查询的时候需要用指针方法
func (dap DnnAddressPool)GetKey()interface{}{
	return dap.Dnn
}

func (dap DnnAddressPool)Less(v1,v2 interface{})bool{
	return v1.(string)<v2.(string)
}

func (dap DnnAddressPool)String()string{
	result:=fmt.Sprintf("Dnn = %s; Address Pool List: ",dap.Dnn)
	for name,_:=range dap.AddressPoolNameList{
		result=result+name+" "
	}
	result=result+"\n"
	
	for tai,v:=range dap.LocationList{
		result=result+fmt.Sprintf("----Tai = %s; Address Pool List: ",tai)
		for name,_:=range v.AddressPoolNameList {
			result = result+name+" "
		}
		result=result+"\n"
	}
	return result+"\n"
}
// implement Entity interface END

func (dap *DnnAddressPool)AddAddressPool(tai string,addrPoolName string){
	if tai !=""{
		if v,ok:=dap.LocationList[tai];ok{
			v.AddressPoolNameList[addrPoolName] = true
		}else {
			lap:=LocationAddressPool{
				Tai:                 tai,
				AddressPoolNameList: make(map[string]bool,0),
			}
			lap.AddressPoolNameList[addrPoolName] = true
			dap.LocationList[tai]=lap
		}
		
		return
	}
	
	dap.AddressPoolNameList[addrPoolName] = true
}

func (dap *DnnAddressPool)DeleteAddressPool(tai string,addrPoolName string){
	// TODO
}

func (dap *DnnAddressPool)GetAddressPoolList(tai string)(poolList []string,err error){
	if tai=="" {
		for k,_:=range dap.AddressPoolNameList{
			poolList=append(poolList,k)
		}
	}else {
		if v,ok:=dap.LocationList[tai];ok{
			for k,_:=range v.AddressPoolNameList{
				poolList=append(poolList,k)
			}
		}
	}
	
	if len(poolList)==0{
		err=errors.New("No Available Address")
	}
	
	return poolList,err
}

func NewDnnAddressPool(dnn string)*DnnAddressPool{
	return &DnnAddressPool{
		Dnn:                 						dnn,
		LocationList:        					make(map[string]LocationAddressPool,0),
		AddressPoolNameList: 		make(map[string]bool,0),
	}
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
