package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// 關鍵對象：
// Client，Server，Request，Response，Header
// 其他輔助的對象：
//

func PrintResponse (resp *http.Response){
	fmt.Println("\nI receive a Response Message.\n\r")
	
	defer resp.Body.Close()
	
	fmt.Println(resp.Proto,resp.Status)
	
	for k,v:=range resp.Header{
		fmt.Println(k,":",v)
	}
	
	fmt.Println("")								// 頭和躰之間必須有一個空行
	rspBody,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatalln(err)
		os.Exit(1)
	}
	if strings.Contains(resp.Header.Get("Content-Type"),"text"){
		fmt.Printf("%s\n\r",rspBody)
		/*
		b:=bytes.NewBuffer([]byte)
		bufio.NewReader(rspBody)
		b.ReadFrom(rspBody)
		rspBody.
		*/
	}else{
		fmt.Println(rspBody)
	}
	
	// Parse time
	genTime,err:=http.ParseTime(resp.Header.Get("Date"))
	if err!=nil{
		log.Fatal(err)
	}else {
		fmt.Println("Escape Time:",time.Since(genTime))
		
	}
}

// 直接調用了缺省client的Post函數發送Post請求
func MyPostRequest(url string, ctxType string){
	body:=bytes.NewBuffer([]byte{})
	body.WriteString("This is my Post Request for "+url)
	resp, err := http.Post(url, ctxType, body)
	if err!=nil{
		log.Fatalln(err)
		return
	}
	
	PrintResponse(resp)
}

func MyGetRequest(url string){
	resp,err:=http.Get(url)
	if err!=nil{
		log.Fatalln(err)
		return
	}
	
	PrintResponse(resp)
}

// 實現自定義的請求的樣例
func CustomizeRequest(method string, pattern string, body io.Reader){
	startTime=time.Now()
	
	client:=&http.Client{}
	
	// Configure customized client
	// TODO
	
	// Generate new request message
	req,err:=http.NewRequest(method,pattern,body)
	if err!=nil{
		log.Fatalln(err)
		return
	}
	
	
	// Set Path Parameters
	// TODO
	
	// Set Query Parameters
	// TODO
	
	// Set Request Header
	// TODO
	
	// Set Header parameters
	// TODO
	
	// Set Request Body
	
	// Call client.Do to send request message, and wait server to response
	rsp, err:=client.Do(req)
	
	if err!=nil{
		// Error process
		log.Fatal(err)
	}else{
		// retrieve response body
		// TODO
		PrintResponse(rsp)
	}
	//
}

func testURLPkg(){
	testURL,err:=url.Parse("https://d00398684:DRC%40748008@localhost:8080/{id}/jerryduren@github/index.jsp?id=5&name=jerry#firstCharter")
	if err!=nil{
		log.Fatal(err)
		return
	}
	
	printURL(testURL)
	
}

// 脱逸码的英文为Escape character。原本是指ASCII中的十进制27，十六进制1B，
// 八进制033所定义的那个字符。对应于标准键盘左上角的ESC键。老式键盘如果没有
// ESC键，替代输入是“Ctrl+[”。在老式的计算机外设控制协议，ASCII码的十进制
// 27这个字符开始的一个字符序列，是外设的控制序列，不能按照这些字符的字面意
// 义解释。后来，就把最初的狭义的Escape character的含义引申开来，在各种计
// 算机语言与协议中，标志着一个转义序列开始的那个字符，都叫做Escape character。
// 最常见的一个例子是C程序设计语言中，用反斜线字符“\”作为脱逸码，来表示那些
// 不可打印的ASCII控制符。在URI协议中，脱逸码是百分号“%”
// 注意URL裏面各部分的關係
// URL = scheme://userInfo(username:password)@host(hostname:port)/URI(path?query#fragement）
func printURL(url *url.URL){
	fmt.Println("--- Print the Structure of URL ---")
	fmt.Println(url)
	fmt.Println("RequestURI: ",url.RequestURI())   // URI是URL裏面去掉path前面的部分，在請求消息裏面帶的uri，是經過逸碼處理后的結果
	fmt.Println("Scheme: ",url.Scheme)
	fmt.Println("Opaque: ",url.Opaque)
	fmt.Println("Username: ",url.User.Username())
	password,_:=url.User.Password()
	fmt.Println("Passwrod: ",password)
	fmt.Println("Host: ",url.Host)
	fmt.Println("Hostname: ",url.Hostname())
	fmt.Println("Port: ",url.Port())
	fmt.Println("Path: ",url.Path)				// 不包括?在内,/{id}/jerryduren@github/index.jsp
	fmt.Println("EscapedPath: ",url.EscapedPath())		// 經過逸碼處理后的path，如：/%7Bid%7D/jerryduren@github/index.jsp
	q:=url.Query()
	fmt.Println("--- Query Parameter ---")
	for k,v:=range q{
		fmt.Println(k,":",v)
	}
	fmt.Println("Fragment: ",url.Fragment)		// 不包括#在内
	fmt.Println("--- END ---")
	
}

var startTime time.Time

func main() {
	
	
	MyPostRequest("http://localhost:8080","application/text")
	
	MyPostRequest("http://localhost:8080/post","application/json")
	
	testURLPkg()
	
	CustomizeRequest("GET","http://localhost:8080/pdu-session/v1?name=jerry&id=398648",nil)
	
	MyGetRequest("http://localhost:8080")
}
