package main

import (
	"net/http"
	"log"
	"github.com/jerryduren/5gc/3gppservice/svcmngserver"
)

func main() {
	
	log.Printf("Server started")
	
	router := svcmngserver.NewRouter()
	
	log.Fatalln(http.ListenAndServe(":8080", router))
	
}
