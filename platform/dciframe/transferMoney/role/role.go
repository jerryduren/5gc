//
package role

import (
	"github.com/pkg/errors"
)

const (
	errInsufficientBalance  = "Insufficient balance"
	errMount = "Mount must be positive"
)

type Accounter interface {
	SetName(string)
	GetName()string
	SetBalance(int)
	GetBalance()int
}


// 定义了一个角色，就是转出账户的源账户
type TransferSource struct {
	Accounter
}

func (t *TransferSource)TransferToTarget(targetAccount Accounter,mount int)error{
	if t.GetBalance()<mount{
		return errors.New(errInsufficientBalance)
	}
	if mount <0{
		return errors.New(errMount)
	}
	
	t.SetBalance(t.GetBalance()-mount)
	targetAccount.SetBalance(targetAccount.GetBalance()+mount)
	
	return nil
}

// 提供角色的构造函数
func NewTransferSource(sourceAccount Accounter)*TransferSource{
	return &TransferSource{sourceAccount}
}