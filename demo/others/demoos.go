package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type Friend struct{
	Name string `json:"name,required"`
	Age int `json:"age,omitempty"`
	Relation string `json:"relation,omitempty"`
}

type Friends []Friend

func (f *Friends)Save(filename string)(err error){
	if f==nil{
		return err
	}
	
	if file,err:=os.Create(filename);err!=nil{
		fmt.Println(err)
		return err
	}else {
		defer file.Close()
		if data,err:=json.MarshalIndent(f,"","  ");err!=nil{
			fmt.Println(err)
			return err
		}else{
			_,err=file.Write(data)
			return err
		}
	}
}

func main() {
	fmt.Println("--- Print Hostname ---")
	hostname,err:=os.Hostname()
	if err!=nil{
		fmt.Println(hostname)
	}else {
		fmt.Println(err)
	}
	
	fmt.Println("--- Print Environment Variables of OS ---")
	for _,v:=range os.Environ(){
		fmt.Println(v)
	}
	
	fmt.Println("--- Print Info of OS ---")
	fmt.Printf("pagesize = %d\n\r",os.Getpagesize())
	if groups,err:=os.Getgroups();err!=nil{
		fmt.Printf("groups = %v\n\r",groups)
	}
	fmt.Printf("egid = %d\n\r",os.Getegid())
	fmt.Printf("uid = %d\n\r",os.Getuid())
	fmt.Printf("euid = %d\n\r",os.Geteuid())
	fmt.Printf("pid = %d\n\r",os.Getpid())
	if wd,err:=os.Getwd();err!=nil{
		fmt.Printf("work dir=%s",wd)
	}else {
		fmt.Println(err)
	}

	fmt.Println("--- 演示如何持久化對象數據 ---")
	friends:=Friends{
		{"Jerry Du",44,"parent"},
		{"Neil",10,"Son"},
		{"Mother",37,""},
		{"Grandfather",50,"Grandfather"},
	}
	if err:=friends.Save(path.Join(os.Getenv("GOBIN"),"friends.json"));err!=nil{
		fmt.Println(err)
	}
	
	fmt.Println("--- 演示如何從JSON文件恢復對象 ---")
	if file,err:=os.Open(path.Join(os.Getenv("GOBIN"),"friends.json"));err!=nil{
		fmt.Println(err)
	}else {
		data:=[]byte{}
		_,err=file.Read(data)
		json.Unmarshal(data,friends)
		fmt.Println(friends)
		file.Close()
	}
}
