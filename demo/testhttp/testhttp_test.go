package testhttp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// GO提供的go test使用示例：
//go test -v 			// 测试目录下的所有文件
//go test -v testhttp.go testhttp_test.go			// 测试制定的文件
//go test -v ./...														// 递归文件夹下的所有文件
//go test -v -run=TestNewTemplate					// 测试指定函数名，函数名可以用通配符，Test*，也可以用正则表达式
//go test -bench=BenchmarkNewTemplate	// 测试只当函数的benchmark
//go test -bench='.' -cpu=3 benchtime=5s		// 指定测试使用的CPU数，测试的时长
//func TestMain(t *testing.M){							// 相当于main函数，在执行测试函数之前可以执行，这样可以初始化以下内容
//
//}
//
//go test -v -coverprofile cover.out user_test.go user.go
//go tool cover -html=cover.out -o cover.html	// 两个命令组合测试代码覆盖率
//
//一些第三方库更可以提供测试支持，尤其httpmock库
//github.com/stretchr/testify
//gihub.com/jarcoal/httpmock


func TestHandleNewUser(t *testing.T) {
	postBody:=url.Values{}
	postBody.Add("say","Hello World!")
	req:=httptest.NewRequest(http.MethodPost,"http://localhost/createuser?name=jerry",strings.NewReader(postBody.Encode()))
	req.Header.Set("Context-Type","application/x-www-form-urlencoded")
	w:=httptest.NewRecorder()
	HandleNewUser(w, req)
	
	if w.Code !=http.StatusOK{
		t.Error("new User API error!")
	}
	
	if w.Body.Len()==0{
		t.Error("response is empty.")
	}
	
	user:=&UserInfo{}
	err:=json.Unmarshal(w.Body.Bytes(),user)
	if err!=nil{
		t.Error("response data error.")
	}
	
	t.Logf("create user api response: %#v",user)
}

func TestHandlePing(t *testing.T) {
	postBody,_:=json.Marshal(Ping{5,"Hello Server!"})
	req:=httptest.NewRequest(http.MethodPost,"http://localhost/ping", strings.NewReader(string(postBody)))
	req.Header.Set("Context-Type","application/json")
	w:=httptest.NewRecorder()
	HandlePing(w, req)
	
	if w.Code !=http.StatusOK{
		t.Error("new User API error!")
	}
	
	if w.Body.Len()==0{
		t.Error("response is empty.")
	}
	
	ping:=&Ping{}
	err:=json.Unmarshal(w.Body.Bytes(),ping)
	if err!=nil{
		t.Error("response data error.")
	}
	
	t.Logf("ping: %#v",ping)
}

func BenchmarkHandlePing(b *testing.B) {
	for i:=0;i<b.N;i++{
		postBody,_:=json.Marshal(Ping{i,"Hello Server!"})
		req:=httptest.NewRequest(http.MethodPost,"http://localhost/ping", strings.NewReader(string(postBody)))
		req.Header.Set("Context-Type","application/json")
		w:=httptest.NewRecorder()
		HandlePing(w, req)
		
		if w.Code !=http.StatusOK{
			b.Error("new User API error!")
		}
		
		if w.Body.Len()==0{
			b.Error("response is empty.")
		}
		
		ping:=&Ping{}
		err:=json.Unmarshal(w.Body.Bytes(),ping)
		if err!=nil{
			b.Error("response data error.")
		}
		
		b.Logf("ping: %#v",ping)
	}
}