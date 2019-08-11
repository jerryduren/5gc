package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

type User struct {
	ID string				`json:"id"`
	Name string 		`json:"name"`
}

type N1Msg struct {
	Apn string
	Imsi string
}

func (a N1Msg)ToByteSlice()([]byte){
	buffer:= bytes.NewBuffer([]byte{})
	buffer.WriteString(a.Apn)
	buffer.WriteString(a.Imsi)

	return buffer.Bytes()
}



func main(){
	users:= []User{{"398648","Jerry Du"},
							  {"335678","Neil Ren"}}
	n1msg:= N1Msg{"cmnet","46000123456789"}

	body := &bytes.Buffer{}
	w:=multipart.NewWriter(body)
	w.SetBoundary("----Boundary")

	// 填part1的消息頭和消息躰
	mHead:=textproto.MIMEHeader{}
	mHead.Add("Content-Type","application/json")
	part1,_:=w.CreatePart(mHead)
	part1Body,_:=json.Marshal(users)
	part1.Write(part1Body)

	// 填part2的消息頭和消息躰
	mHead = textproto.MIMEHeader{}
	mHead.Add("Context-Type","application/vnd.3gpp.5gnas")
	mHead.Add("Content-Id"," n1msg")
	part2,_:=w.CreatePart(mHead)
	part2.Write(n1msg.ToByteSlice())

	w.Close()

	fmt.Println(body.String())
}

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

func ApiBasePathHandlerFunc(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"This is API base path, please access right API path.")
}


func PostSmContexts(w http.ResponseWriter,r *http.Request){
	// 下面的代碼演示了如何從multipart/related的MIME類型中讀取數據
	// Read multipart/related body
	_,params,_:=mime.ParseMediaType(r.Header.Get("Content-Type"))
	mr:=multipart.NewReader(r.Body,params["boundary"])
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}
		
		slurp, err := ioutil.ReadAll(p)		//  讀的是part的body部分，不包括part的head部分
		if err != nil {
			log.Fatal(err)
			break
		}
		
		locData:=model.SmContextCreateData{}
		if strings.Contains(p.Header.Get("Content-Type"),"application/json"){
			err:=json.Unmarshal(slurp,&locData)
			if err!=nil{
				log.Fatal()
				break
			}
			// TODO
			fmt.Println("jsondata:\n",locData)
		}
		
		if strings.Contains(p.Header.Get("Content-Type"),"application/vnd.3gpp.5gnas"){
			fmt.Println("n1msg:\n",slurp)
		}
	}
	
	io.WriteString(w,"This is PostSmContexts.")
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