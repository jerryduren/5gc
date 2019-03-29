package nnrfnfmanagement

import (
	"github.com/jerryduren/5gc/3gppservice/svcmngclient"
	"golang.org/x/net/context"
)

var nrfMngClient *svcmngclient.APIClient

func init(){
	cfg:= svcmngclient.NewConfiguration()
	cfg.BasePath="http://localhost:8080/nnrf-nfm/v1"
	nrfMngClient=svcmngclient.NewAPIClient(cfg)
}

func NewNfProfile()*svcmngclient.NfProfile{
	
	
	return &svcmngclient.NfProfile{}
}
func RegisterNFInstance(nfInstanceID svcmngclient.instance)  {
	nfProfile,rsp,err:=nrfMngClient.NFInstanceIDDocumentApi.RegisterNFInstance(context.Background(),NewNfProfile(),nfInstanceID)
	
}

func UpdateNFInstance(){

}
