package main

import (
	"fmt"
	
	"github.com/jerryduren/5gc/3gppservice/models"
	 "github.com/jerryduren/5gc/nfs/nrf/nfprofileentity"
)

func main() {
	
	// 结构体可以直接赋值
	nfProfile1:=models.NewNfProfile(models.NextNfInstanceID(),models.NF_TYPE_AMF,models.NF_STATUS_REGISTERED)
	nfProfile1.PlmnList = []models.PlmnId{models.PlmnId{"460", "00"}, models.PlmnId{"460", "02"}}
	nfProfile1.UdmInfo = models.UdmInfo{
		GroupId:                        "223",
		SupiRanges:                     nil,
		GpsiRanges:                     nil,
		ExternalGroupIdentifiersRanges: nil,
		RoutingIndicators:              nil,
	}
	nfProfile1.UdmInfo.GpsiRanges = []models.IdentityRange{
		models.IdentityRange{
			Start:   "13917290000",
			End:     "13917299999",
			Pattern: "1124",
		},
		models.IdentityRange{
			Start:   "18621070000",
			End:     "18621079990",
			Pattern: "1124",
		},
	}
	
	nfProfile2:= nfprofile.NewNfProfileEntityFromNfProfile(*nfProfile1)
	
	fmt.Println(*nfProfile1)
	fmt.Println(nfProfile2)
	
	// GO语言里面结构体直接赋值很危险，两次打印的nfProfile3的值是不一样的，
	// 原因是nfProfile1修改了结构体内层的切片的值。
}
