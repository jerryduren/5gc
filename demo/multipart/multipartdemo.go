package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/textproto"

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


