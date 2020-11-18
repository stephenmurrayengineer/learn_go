package main

type Wallet struct {
	balance int	
}


// when referencing Wallet in these functions, a copy is made.
// Therefore, use a pointer
func (w *Wallet) Deposit(amount int) {
	// notice that you don't need to dereference the pointer:
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}