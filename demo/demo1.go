package main

import (
	"encoding/json"
	"fmt"
)

type Demo struct {
	F1 string			`json:"f_1"`
	F2 string			`json:"f_2"`
}

type Demo1 struct {
	Demo
}

type Demo2 struct {
	Demo Demo				`json:"demo"`
}

func main(){
	var demo1 Demo1
	var demo2 Demo2
	
	demo1.Demo.F1 = "f1"
	demo1.Demo.F2 = "f2"
	demo1.F1 = "f1"
	demo1.F2 = "f2"
	
	demo2.Demo.F1 = "f1"
	demo2.Demo.F2 = "f2"
	
	data1,_:=json.Marshal(demo1)
	data2,_:=json.Marshal(demo2)
	fmt.Println(string(data1))
	fmt.Println(string(data2))
	
	var demo Demo
	demo = demo1
	
}
