package wallet

import (
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("balance of wallet: ", w.balance)

	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	fmt.Println("balance of wallet: ", w.balance)

	w.balance -= amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
