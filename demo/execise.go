package main

import (
	"fmt"
	"time"
	
	"github.com/jerryduren/5gc/prp"
	"github.com/jerryduren/5gc/smf"
)

func main() {
	
	// test UUID
	fmt.Println(prp.GenerateUUID())
	
	fmt.Println(prp.GenerateUUID())
	fmt.Println(prp.GenerateUUID())
	fmt.Println(prp.GenerateUUID())
	
	startTime:=time.Now()
	for i:=0;i<1000;i++{
		fmt.Println(smf.NewSupi(460,00))
	}
	fmt.Println("escape time:",time.Since(startTime))
}
