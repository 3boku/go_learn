package accounts

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