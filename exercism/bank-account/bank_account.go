package account

import "sync"

type Account struct {
	sync.RWMutex
	open    bool
	balance int64
}

func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}
	return &Account{open: true, balance: deposit}
}

func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if !account.open || account.balance+amount < 0 {
		return 0, false
	}
	account.balance += amount
	return account.balance, true
}

func (account *Account) Balance() (balance int64, ok bool) {
	account.RLock()
	defer account.RUnlock()
	if !account.open {
		return 0, false
	}
	return account.balance, true
}

func (account *Account) Close() (payout int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if !account.open {
		return 0, false
	}
	account.open = false
	return account.balance, true
}
