package main

import (
	"fmt"
	"name/accounts"
)



func main(){
	account := accounts.NewAccount("sambok")
	fmt.Println(account)
}