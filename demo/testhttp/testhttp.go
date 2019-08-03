package testhttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type UserInfo struct {
	Name string	`json:"name"`
	Age  int			`json:"age"`
}

func NewUserInfo(name string)*UserInfo{
	return &UserInfo{Name:name,Age:rand.Int()}
}

func HandleNewUser(w http.ResponseWriter,r *http.Request){
	name:=r.URL.Query().Get("name")
	fmt.Println("URL parameter user name is ",name)
	say:=r.FormValue("say")
	fmt.Println("req say: ",say)
	newUser:=NewUserInfo(name)
	jData,_:=json.Marshal(newUser)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type","application/json")
	w.Write(jData)
}

type Ping struct {
	Sequence int				`json:"sequence"`
	Context string				`json:"context"`
}

func HandlePing(w http.ResponseWriter,r *http.Request){
	//// 读取请求对象的json body
	//buffer := bytes.NewBuffer(make([]byte, 4096))
	//_, err := io.Copy(buffer, r.Body)
	//if err !=nil{
	//	// TODO
	//}
	
	ping:=&Ping{}
	
	//// method 1
	//if strings.Contains(r.Header.Get("Context-Type"), "json"){
	//	json.NewDecoder(r.Body).Decode(ping)
	//}
	
	// method 2
	data,_:=ioutil.ReadAll(r.Body)
	json.Unmarshal(data,ping)
	
	jData,_:=json.Marshal(ping)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}