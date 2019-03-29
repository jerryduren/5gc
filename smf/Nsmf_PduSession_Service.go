package smf

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/jerryduren/5gc/prp"
)

var(
	NsmfPduSessionServer *http.Server

)

func init(){
	// init all kinds of variable
	// TODO
	
	// 定義自己的路由器，並注冊處理函數
	mux:=http.NewServeMux()
	mux.HandleFunc("/nsmf-pdusession/v1",ApiBasePathHandlerFunc)
	mux.HandleFunc("/nsmf-pdusession/v1/sm-contexts",PostSmContexts)
	mux.HandleFunc("/nsmf-pdusession/v1/sm-contexts/{smContextRef}/retrieve",RetrieveSmContext)
	mux.HandleFunc("/nsmf-pdusession/v1/sm-contexts/sm-contexts/{smContextRef}/modify",UpdateSmContext)
	mux.HandleFunc("/nsmf-pdusession/v1/sm-contexts/{smContextRef}/release",ReleaseSmContext)
	mux.HandleFunc("/nsmf-pdusession/v1/sm-contexts/pdu-sessions",PostPduSessions)
	mux.HandleFunc("/nsmf-pdusession/v1/pdu-sessions/{pduSessionRef}/modify",UpdatePduSession)
	mux.HandleFunc("/nsmf-pdusession/v1/pdu-sessions/{pduSessionRef}/release",ReleasePduSession)
	
	// 初始化Server
	NsmfPduSessionServer = &http.Server{
		Addr:           ":8080",
		Handler:		mux,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// TODO
	}
	
	// 注冊shutdown的時候的處理函數，測試結果好像沒有調用啊！！！
	NsmfPduSessionServer.RegisterOnShutdown(onShutdownNsmfPduSessionServer)
	
	
}

func onShutdownNsmfPduSessionServer(){
	fmt.Println("After 2 seconds, I will shutdown, please save all modifies!")
	time.Sleep(time.Second*2)
}

func StartNsmfPduSessionService() error {
	prp.DebugInfo("Nsmf_Pdu_Session Service is started!")
	if err:=NsmfPduSessionServer.ListenAndServe();err!=nil{
		prp.DebugInfo("Start Nsmf_Pdu_Session Failure!")
		return err
	}
	
	return nil
}

func ShutdownNsmfPduSessionService()error{
	prp.DebugInfo("I will close Nsmf_Pdu_Session Service!")
	return NsmfPduSessionServer.Close()
}
