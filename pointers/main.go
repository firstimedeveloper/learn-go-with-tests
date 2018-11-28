package main

import (
	"fmt"
	"github.com/firstimedeveloper/learn-go-with-tests/pointers/wallet"
)

func main() {
	w := Wallet{Bitcoin(100)}
	fmt.Println(w.Balance())
}
