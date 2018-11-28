package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with sufficient funds", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(10)

		want := Bitcoin(0)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	//t.Helper()
	got := wallet.Balance()

	fmt.Println("balance of wallet: ", &wallet.balance)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	//t.Helper()
	if got != nil {
		t.Fatal("got an error but ddn't want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	//t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', wanted '%s'", got, want)
	}
}
