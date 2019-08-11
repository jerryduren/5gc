package data

type Account struct {
	name string
	balance int
}

func (a *Account)SetName(name string){
	if len(name)!=0 {
		a.name = name
	}
}

func (a *Account)GetName()string{
	return a.name
}

func (a *Account)SetBalance(mount int){
	if mount > 0 {
		a.balance = mount
	}
}

func (a *Account)GetBalance()int{
	return a.balance
}

func NewAccount(name string, balance int)*Account{
	return &Account{name,balance}
}