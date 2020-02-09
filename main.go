package main

import (
	"fmt"

	"github.com/sehoon123/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
		// log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
