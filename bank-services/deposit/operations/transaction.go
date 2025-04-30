package operations

import "fmt"

// TODO возможно стоит переименовать в банковскую транзакцию
type Transaction struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Type        string  `json:"type"`
}

var Transactions []Transaction

func CreateTransaction(id, fromAccount, toAccount string, amount float64, transactionType string) Transaction {
	tx := Transaction{
		ID:          id,
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Amount:      amount,
		Type:        transactionType,
	}
	Transactions = append(Transactions, tx)
	return tx
}

func GetAllTransactions() []Transaction {
	return Transactions
}

func Transfer(fromAccount, toAccount string, amount float64) (Transaction, error) {
	if amount <= 0 {
		return Transaction{}, fmt.Errorf("amount can't be zero or lower")
	}
	tx := CreateTransaction(fmt.Sprintf("%d", len(Transactions)+1), fromAccount, toAccount, amount, "transfer")
	return tx, nil
}
