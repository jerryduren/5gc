package smf

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
	
	"github.com/jerryduren/5gc/model"
)

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

func RetrieveSmContext(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"This is RetrieveSmContext.")
}

func UpdateSmContext(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"This is UpdateSmContext.")
}

func ReleaseSmContext(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"This is ReleaseSmContext.")
}

func PostPduSessions(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"This is PostPduSessions.")
}

func UpdatePduSession(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"This is UpdatePduSession.")
}

func ReleasePduSession(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"This is ReleasePduSession.")
}
