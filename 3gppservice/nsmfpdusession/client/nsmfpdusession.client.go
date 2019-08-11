package main

import (
	"context"
	"fmt"
	
	"github.com/antihax/optional"
	"github.com/jerryduren/5gc/3gppservice/nsmfpdusession/client/client"
)

func main(){
	cfg:=client.Configuration{
		BasePath:      "/nsmf-pdusession/v1",
		Host:          "http://127.0.0.1:8080",
		Scheme:        "",
		DefaultHeader: make(map[string]string),
		UserAgent:     "AMF v1",
	}
	pduSessionClient:=client.NewAPIClient(&cfg)
	
	smContextCreateData:=client.SmContextCreateData{
		Supi:                    "460051234567890",
		UnauthenticatedSupi:     false,
		Pei:                     nil,
		Gpsi:                    nil,
		PduSessionId:            nil,
		Dnn:                     nil,
		SNssai:                  nil,
		HplmnSnssai:             nil,
		ServingNfId:             nil,
		Guami:                   nil,
		ServiceName:             nil,
		ServingNetwork:          nil,
		RequestType:             client.RequestType{},
		N1SmMsg:                 nil,
		AnType:                  nil,
		SecondAnType:            nil,
		RatType:                 nil,
		PresenceInLadn:          nil,
		UeLocation:              nil,
		UeTimeZone:              nil,
		AddUeLocation:           nil,
		SmContextStatusUri:      nil,
		HSmfUri:                 nil,
		AdditionalHsmfUri:       nil,
		OldPduSessionId:         nil,
		PduSessionsActivateList: nil,
		UeEpsPdnConnection:      "",
		HoState:                 client.HoState{},
		PcfId:                   nil,
		NrfUri:                  nil,
		SupportedFeatures:       nil,
		SelMode:                 client.DnnSelectionMode{},
		BackupAmfInfo:           nil,
		TraceData:               nil,
		UdmGroupId:              nil,
		RoutingIndicator:        "",
		EpsInterworkingInd:      client.EpsInterworkingIndication{},
		IndirectForwardingFlag:  false,
		TargetId:                nil,
		EpsBearerCtxStatus:      "",
		CpCiotEnabled:           false,
		InvokeNef:               false,
		MaPduIndication:         false,
		N2SmInfo:                nil,
		SmContextRef:            nil,
	}
	reqBody:=client.PostSmContextsOpts{optional.NewInterface(smContextCreateData)}
	rspBody,_, err:=pduSessionClient.SMContextsCollectionApi.PostSmContexts(context.Background(),&reqBody)
	if err!=nil{
		fmt.Println("Response Failure")
	}else{
		fmt.Print(rspBody)
	}
}
