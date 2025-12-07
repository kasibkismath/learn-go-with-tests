package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Points(10.00))

	got := wallet.Balance()
	want := Points(10.00)

	if got != want {
		// when %s is used, it implicitly calls the String() method
		t.Errorf("got %s want %s", got, want)
	}
}
