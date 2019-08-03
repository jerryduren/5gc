package main

import (
	"fmt"
	
	"github.com/jerryduren/ngc/plt/lib/repository"
)

func main() {
		var 	pduSessionRepository repository.Repository
		pduSessionRepository = repository.NewPduSessionEntityRepository()
		
		pduSessionRepository.Add(repository.PduSession{"jerry","cmnet" })
	    pduSessionRepository.Add(repository.PduSession{"neil","cmwap"})
		
		cnt:= pduSessionRepository.Count()
		if cnt!=2{
			fmt.Println("ERROR")
		}else {
			fmt.Println("Has ",cnt, "records")
		}
}

