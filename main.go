package main

import (
	"fmt"

	"github.com/sehoon123/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
