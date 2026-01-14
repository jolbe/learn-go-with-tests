// Package arrays showcases how to use arrays with generics
package arrays

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{from.Name, to.Name, sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(account Account, transaction Transaction) Account {
	if transaction.From == account.Name {
		account.Balance -= transaction.Sum
	}
	if transaction.To == account.Name {
		account.Balance += transaction.Sum
	}
	return account
}
