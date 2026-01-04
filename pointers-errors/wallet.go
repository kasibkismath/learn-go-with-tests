package pointerserrors

import (
	"errors"
	"fmt"
)

type Points float64

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Points
}

func (w *Wallet) Deposit(amount Points) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Points) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w Wallet) Balance() Points {
	return w.balance
}

// implicitly matches the Stringer interface
func (p Points) String() string {
	return fmt.Sprintf("%2f PTS", p)
}
