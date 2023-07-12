package main

import (
	"fmt"
	"name/accounts"
)



func main(){
	account := accounts.NewAccount("sambok")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())

}