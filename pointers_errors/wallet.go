package main

import "fmt"

type Bitcoin int

// function to return Bitcoin in parsed string:
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// when referencing Wallet in these functions, a copy is made.
// Therefore, use a pointer
func (w *Wallet) Deposit(amount Bitcoin) {
	// notice that you don't need to dereference the pointer:
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}