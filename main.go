package main

import (
	"fmt"

	"github.com/sehoon123/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
}
