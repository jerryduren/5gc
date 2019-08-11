package main

import (
	"fmt"
	"os"
	
	"github.com/jerryduren/5gc/platform/dciframe/transferMoney/context"
	"github.com/jerryduren/5gc/platform/dciframe/transferMoney/data"
)

func main() {
	sourceAccount:=data.NewAccount("Jerry Du",1000)
	targetAccount:=data.NewAccount("Neil Ren",300)
	if err:=context.TransferMoney(sourceAccount, targetAccount,-400);err!=nil{
		fmt.Println("Transfer money failure:",err)
		os.Exit(1)
	}
	
	fmt.Println(sourceAccount.GetName(), ": ",sourceAccount.GetBalance())
	fmt.Println(targetAccount.GetName(), ": ", targetAccount.GetBalance())
}
