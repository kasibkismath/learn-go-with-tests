package pointerserrors

import "fmt"

type Points float64

type Wallet struct {
	balance Points
}

func (w *Wallet) Deposit(amount Points) {
	w.balance += amount
}

func (w Wallet) Balance() Points {
	return w.balance
}

// implicitly matches the Stringer interface
func (p Points) String() string {
	return fmt.Sprintf("%2f PTS", p)
}
