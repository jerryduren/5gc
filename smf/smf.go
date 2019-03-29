package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/jerryduren/5gc/3gppservice/ts29502nsmfpduserver"
)

func main() {
	
	// Start  Nsmf_PDUSession Service
	fmt.Println("Nsmf_PDUSession service is running ...")
	log.Println("Nsmf_PDUSession service is started.")
	NsmfPduSessionRouter := ts29502nsmfpduserver.NewRouter()
	
	log.Fatalln(http.ListenAndServe(":8080",NsmfPduSessionRouter))
	
}
