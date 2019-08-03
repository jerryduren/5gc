package repository

import (
	"fmt"
)

type PduSession struct {
	Id uint8
	Dnn string
}
func (p *PduSession)String()string{
	return fmt.Sprintf("{%d: %s}",p.Id,p.Dnn)
}
func (p PduSession)GetKey()interface{}{
	return p.Id
}

func (p PduSession)Less(v1,v2 interface{})bool{
	return v1.(uint8)<v2.(uint8)							// v1,v2是指的p.GetKey得到的类型的比较方法
}

func DemoPduSessionRepository() {
	pduSessionRepository := NewEntityRepository(&PduSession{})
	
	_ = pduSessionRepository.Add(&PduSession{5, "cmnet"})
	_= pduSessionRepository.Add(&PduSession{Id: 6, Dnn: "cmwap"})
	
	fmt.Println(pduSessionRepository)
	
}
