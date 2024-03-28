package twg

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got deposit %v, want %v \n", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("withdraw insufficient costs", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(30)

		if err == nil {
			t.Errorf("Wanted error but didn't receive one")
		}
	})
}
