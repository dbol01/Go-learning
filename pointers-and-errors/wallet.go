package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin type for clarity
type Bitcoin int

// Stringer -
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet - holds private fields/methods
type Wallet struct {
	balance Bitcoin
}

// Balance method on wallet to allow public access to balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit method on wallet to allow public deposits
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds - error message
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw method on wallet to allow public withdrawals
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
