package accounts

import "errors"

// Account : Whatever
type Account struct {
	owner   string
	balance int
}

var errMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw from your Account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errMoney
	}
	a.balance -= amount
	return nil
}
