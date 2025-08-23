package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want error) {
		if got == nil {
			t.Errorf("wanted an error but didn't get one")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		// It's important to check that no money was removed from the account
		// in illegitimate scenarios! In case of withdrawal with insufficient
		// funds, the final balance should be the same as the starting balance
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}
