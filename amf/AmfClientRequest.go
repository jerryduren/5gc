package amf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	
	"github.com/jerryduren/5gc/model"
	"github.com/jerryduren/5gc/plt"
)

func HeartBeatBetweenAmfAndSmf()error{
	_,err:=NsmfPduSessionClient.Get(SmfBaseURL)
	if err!=nil{
		log.Fatal(err)
		plt.ErrorInfo("Heartbeat is failure between AMF and SMF.")
		return err
	}else {
		fmt.Println("Path test is succeed between AMF and SMF.")
		return nil
	}
}

// For test, new a SMContextCreateBody
func NewSMContextCreateBody()*model.SmContextCreateBody{
	body:=model.SmContextCreateBody{
		JsonData:model.SmContextCreateData{},
		BinaryDataN1SmMessage:make([]byte,0),
	}
	body.JsonData.ServingNfId = "SMFID"
	body.JsonData.Supi = "46000123456789"
	body.JsonData.AnType = "3GPP_ACCESS"
	body.JsonData.AdditionalHsmfUri = []string{"http://localhost:8090"}
	body.JsonData.SmContextStatusUri = "/ddd/eee"
	body.BinaryDataN1SmMessage = []byte{48,49,50,51}
	
	return &body
}


func PostSmContexts(body *model.SmContextCreateBody) (model.SmContextCreatedData, *http.Response, error) {
	/*
	mydata:=model.SmContextCreateData{}
	mydata.Supi = "46000123456789"
	mydata.HSmfUri = "http://localhost:8080/"
	mydata.AdditionalHsmfUri = []string{"https://localhost:9090"}
	
	marshelData,_:=json.Marshal(mydata)
	
	ioR:=bytes.NewReader(marshelData)
	NsmfPduSessionClient.Post(SmfBaseURL+"/sm-contexts","application/json",ioR)
	*/
	
	
	localVarReturnValue :=  model.SmContextCreatedData{}
	
	// PostSmContexts操作的消息躰是一個multipart/related的數據類型
	// 由JSON數據部分+N1 ANS.1二進制消息兩部分組成
	localBody:=bytes.NewBuffer([]byte{})
	w:=multipart.NewWriter(localBody)
	w.SetBoundary(Boundary)
	// Set JSON part
	mHead:=make(textproto.MIMEHeader)
	mHead.Set("Content-Type","application/json")
	part,err:=w.CreatePart(mHead)
	if err!=nil{
		log.Fatal(err)
		plt.ErrorInfo("Create PostSmContexts Body Error!")
		return localVarReturnValue,nil,err
	}
	partData,err:=json.Marshal(body.JsonData)
	if err!=nil{
		log.Fatal(err)
		plt.ErrorInfo("Json marshal failure for PostSmContexts!")
		return localVarReturnValue,nil,err
	}
	_,err=part.Write(partData)
	if err!=nil{
		log.Fatal(err)
		return localVarReturnValue,nil,err
	}
	// Set N1MSG ANS.1
	mHead.Set("Content-Type","application/vnd.3gpp.5gnas")
	mHead.Set("Content-Id","n1msg")
	part,err=w.CreatePart(mHead)
	if err!=nil{
		log.Fatal(err)
		return localVarReturnValue,nil,err
	}
	_,err=part.Write(body.BinaryDataN1SmMessage)
	if err!=nil{
		log.Fatal(err)
		return localVarReturnValue,nil,err
	}
	w.Close()
	
	fmt.Println("localBody:",localBody)
	
	// Init new request message
	req, err:=http.NewRequest("POST",SmfBaseURL+"/sm-contexts",bytes.NewReader(localBody.Bytes()))
	if err!=nil{
		log.Fatal(err)
		plt.ErrorInfo("Create Request is failure for PostSmContexts.")
		return localVarReturnValue,nil,err
		
	}
	
	// Set Path Parameters
	// TODO
	
	// Set Query Parameters
	// TODO
	
	// Set Request Header
	// TODO
	
	// Set Header parameters
	req.Header.Set("Content-Type","multipart/related;boundary=----Boundary")
	
	fmt.Print(req)
	
	rsp,err:=NsmfPduSessionClient.Do(req)
	if err!=nil{
		plt.ErrorInfo("Initiate PostSmContexts failure.")
		log.Fatal(err)
		// TODO
		return model.SmContextCreatedData{},nil,nil
	}else{
		fmt.Println(rsp)
	}
	
	return model.SmContextCreatedData{},rsp,nil
}


func RetrieveSmContext(body model.SmContextCreateBody) (model.SmContextCreatedData, *http.Response, error) {
	
	return model.SmContextCreatedData{},nil,nil
}

func UpdateSmContext(body model.SmContextCreateBody) (model.SmContextCreatedData, *http.Response, error) {
	return model.SmContextCreatedData{},nil,nil
}

func ReleaseSmContext(body model.SmContextCreateBody) (model.SmContextCreatedData, *http.Response, error) {
	return model.SmContextCreatedData{},nil,nil
}

func PostPduSessions(body model.SmContextCreateBody) (model.SmContextCreatedData, *http.Response, error) {
	return model.SmContextCreatedData{},nil,nil
}

func UpdatePduSession(body model.SmContextCreateBody) (model.SmContextCreatedData, *http.Response, error) {
	return model.SmContextCreatedData{},nil,nil
}

func ReleasePduSession(body model.SmContextCreateBody) (model.SmContextCreatedData, *http.Response, error) {
	return model.SmContextCreatedData{},nil,nil
}

