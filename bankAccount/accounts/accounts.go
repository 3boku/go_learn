package accounts

import "errors"


var errNoMoney = errors.New("can't withdraw")
//account struct
type Account struct{
	owner string
	balance int
}

//New Account creates a new Account
func NewAccount(owner string) * Account{
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int){
	a.balance += amount
}

//Balance x amount on your account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount{
		return errNoMoney
	}
	a.balance -= amount
	return nil
}