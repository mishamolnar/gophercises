package twg

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(num Bitcoin) {
	w.balance += num
}

func (w *Wallet) Withdraw(num Bitcoin) error {
	if num >= w.balance {
		return fmt.Errorf("Could not ")
	}
	w.balance -= num
	return nil
}
