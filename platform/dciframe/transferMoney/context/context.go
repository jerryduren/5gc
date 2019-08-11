// Context的作用是创建Role，并且把领域对象与角色绑定起来
package context

import "github.com/jerryduren/5gc/platform/dciframe/transferMoney/role"

func TransferMoney(sourceAccount, targetAccount role.Accounter, mount int)error{
	return  role.NewTransferSource(sourceAccount).TransferToTarget(targetAccount,mount)
}
