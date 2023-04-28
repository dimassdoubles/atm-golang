package transaction

import (
)

type TransactionsRepository struct {
	Data []Transaction
}

func (tr *TransactionsRepository) Add(transaction Transaction) {
	tr.Data = append(tr.Data, transaction)
}

func (tr TransactionsRepository) FilterByCardNumber(cardNumber string) []Transaction {
	result := []Transaction{}

	for i:= range tr.Data {
		if tr.Data[i].GetCardNumber() == cardNumber {
			result = append(result, tr.Data[i])
		}
	}

	return result
}