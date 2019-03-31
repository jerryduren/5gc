package main

import (
	"fmt"
	"log"

	smfclient "github.com/jerryduren/5gc/smf/TS29502_Nsmf_PDUSession/client"
	"golang.org/x/net/context"
)
func main() {
	cfg:=smfclient.NewConfiguration()
	cfg.UserAgent = "SMF Client for test SMF Server"
	cfg.Host = "http://127.0.0.1:8080"
	cfg.BasePath = "nsmf-pdusession/v1"

	smfClient:=smfclient.NewAPIClient(cfg)

	body:=smfclient.NewBody()

	smContextRspDate,rsp,err:=smfClient.SMContextsCollectionApi.PostSmContexts(context.Background(),*body)
	if err!=nil{
		log.Fatalln("Create Pdu Session Failure!\n",err)
	}else {
		fmt.Println(rsp.Location())
		fmt.Println(smContextRspDate.PduSessionId)
		//
	}
}

