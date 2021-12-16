package account

// Define the Account type here.
type Account struct {
	amt    int64
	status bool
}

func Open(amount int64) *Account {
	var acc *Account
	if amount < 0 {
		return nil
	}
	acc.amt = amount
	acc.status = true
	return acc
	// panic("Please implement the Open function")
}

func (a *Account) Balance() (int64, bool) {
	if a.amt > 0 && a.status {
		return a.amt, true
	} else {
		return 0, false
	}
	// panic("Please implement the Balance function")
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	return a.amt + amount, true
	// panic("Please implement the Deposit function")
}

func (a *Account) Close() (int64, bool) {
	if a.status {
		a.status = false
	}
	a.amt = 0
	return 0, false
	// panic("Please implement the Close function")
}
