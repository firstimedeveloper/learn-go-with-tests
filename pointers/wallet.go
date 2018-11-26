package wallet

import (
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("adr of balance in test: ", &w.balance)

	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
