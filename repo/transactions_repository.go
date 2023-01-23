package repo

import (
	"git.solusiteknologi.co.id/golab/atm/transaction"
)

type TransactionsRepository struct {
	Data []transaction.Transaction
}

func (tr *TransactionsRepository) Add(transaction transaction.Transaction) {
	tr.Data = append(tr.Data, transaction)
}

func (tr TransactionsRepository) FilterByCardNumber(cardNumber string) []transaction.Transaction {
	result := []transaction.Transaction{}

	for i:= range tr.Data {
		if tr.Data[i].GetCardNumber() == cardNumber {
			result = append(result, tr.Data[i])
		}
	}

	return result
}