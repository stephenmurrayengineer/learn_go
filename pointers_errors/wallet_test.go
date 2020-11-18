package main

import "testing"
// import "fmt"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
	
		got := wallet.Balance()
		want := Bitcoin(10)
	
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		wallet.Withdraw(Bitcoin(10))
	
		got := wallet.Balance()
		want := Bitcoin(10)
	
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}