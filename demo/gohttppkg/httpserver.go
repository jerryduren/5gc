package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)



func MyPostHandler(w http.ResponseWriter,r *http.Request){
	PrintRequestMessage(r)
	fmt.Fprintf(w,"%s","This is my first handlerfunc!")
}

func PrintRequestMessage(r *http.Request){
	
	URL,_:=url.ParseRequestURI(r.RequestURI)
	fmt.Println("\nI reveive a Request Message.",URL)
	
	fmt.Printf("%s %s %s\n",r.Method,r.RequestURI,r.Proto)	// 請求行，URI是URL除path之前的部分
	
	fmt.Println("Host:",r.Host)
	for k,v:=range r.Header{
		fmt.Println(k,":",v)
	}
	
	fmt.Println("")				// 請求行和躰之間必須有一個空行
	reqBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Fatal(err)
		return
	}else {
		fmt.Println(reqBody)
	}
}

type MyResource struct {

}

func (a *MyResource)ServeHTTP (w http.ResponseWriter,r *http.Request){
	PrintRequestMessage(r)
	
	fmt.Fprintf(w,"%s","This is post response of MyResource!")
}

func main(){
	http.Handle("/",&MyResource{})				// 不考慮用的什麽方法
	http.HandleFunc("/post",MyPostHandler)		// 不考慮用的什麽方法
	http.HandleFunc("/pdu-session/v1",func(w http.ResponseWriter,r *http.Request){
													PrintRequestMessage(r)
													fmt.Fprintf(w,"%s\n\r","This is /pud-session handler.")
													if r.ParseForm()!=nil{
														log.Fatal("Parse Form error!")
														return
													}
													// may read form info.
													fmt.Println("--- Print Form Info ---")
													for k,v:=range r.Form{
														fmt.Println(k,v)
													}
												})
	
	fmt.Println("Server is started in 8080 port!")
	log.Fatal(http.ListenAndServe(":8080",nil))
}