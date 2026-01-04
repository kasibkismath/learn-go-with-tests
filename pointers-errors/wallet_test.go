package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Points(10.00))
		assertBalance(t, wallet, Points(10.00))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Points(20.00)}
		err := wallet.Withdraw(Points(10.00))
		assertNoError(t, err)
		assertBalance(t, wallet, Points(10.00))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Points(20.00)}
		err := wallet.Withdraw(Points(100.00))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Points(20.00))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Points) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
