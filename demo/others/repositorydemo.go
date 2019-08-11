package main

import (
	"fmt"
	
	"github.com/jerryduren/5gc/platform/lib/repository"
)

func main() {
		var 	pduSessionRepository repository.Repository
		pduSessionRepository = repository.NewEntityRepository(&repository.PduSession{})
		
		// 这个地方添加的时候类型一定要与PduSession的类型一致，不能直接填5
		pduSessionRepository.Add(&repository.PduSession{uint8(5),"cmnet" })
	    pduSessionRepository.Add(&repository.PduSession{uint8(6),"cmwap"})
		
		cnt:= pduSessionRepository.Count()
		if cnt!=2{
			fmt.Println("ERROR")
		}else {
			fmt.Println("Has ",cnt, "records")
		}
		
		pduSession,_:=pduSessionRepository.FindById(uint8(5))
		fmt.Println("Key: ", pduSession.GetKey())
		fmt.Println("Pdu Session: ", pduSession)
		
		// 有两点要千万注意：实现Entity的时候不能用指针方法！类型查询的时候是对象的指针类型！
		v:=pduSession.(*repository.PduSession)
		fmt.Println(v.Dnn)
		v.Dnn = "another"
		
		// 上面修改了v.Dnn不会在库里面生效
		pduSession,_=pduSessionRepository.FindById(uint8(5))
		fmt.Println("Key: ", pduSession.GetKey())
		fmt.Println("Pdu Session: ", pduSession)
	
		// 需要重新Add操作才会生效
		pduSessionRepository.Add(v)
	
		pduSession,_=pduSessionRepository.FindById(uint8(5))
		fmt.Println("Key: ", pduSession.GetKey())
		fmt.Println("Pdu Session: ", pduSession)
		
}

