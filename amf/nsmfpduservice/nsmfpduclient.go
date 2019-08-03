package nsmfpduservice

import (
	"github.com/jerryduren/5gc/3gppservice/nsmfpduclient"
	"golang.org/x/net/context"
)

var nsmfClientCfg *nsmfpduclient.Configuration
var nsmfClientSvc *nsmfpduclient.APIClient


func init (){
	nsmfClientCfg =nsmfpduclient.NewConfiguration()
	nsmfClientCfg.BasePath = "http://localhost:8080//nsmf-pdusession/v1"
	
	nsmfClientSvc=nsmfpduclient.NewAPIClient(nsmfClientCfg)
}

func NewPduSessionCreateData() *nsmfpduclient.PduSessionCreateData{
	localPduSessionCreateData:=nsmfpduclient.PduSessionCreateData{}
	
	localPduSessionCreateData.
	
	return &localPduSessionCreateData
}
func CreatePDUSession(){
	pduSessionCreatedData,rsp,err:=nsmfClientSvc.PDUSessionsCollectionApi.PostPduSessions(context.Background(),pduSessionCreateData)
}
