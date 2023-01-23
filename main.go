package main

import (
	"fmt"
	"git.solusiteknologi.co.id/golab/atm/card"
	"git.solusiteknologi.co.id/golab/atm/transaction"
	"git.solusiteknologi.co.id/golab/atm/repo"
)

func main() {

	transactions := repo.TransactionsRepository{
		Data : []transaction.Transaction{},
	}

	cards := repo.CardsRepository{
		Data : []card.Card{
			&card.EMoney{
				CardInfo: card.CardInfo{
					CardNumber: "9999999999",
					Pin: "999999",
					Balance: 100,
				},
			},
			&card.ATM{
				CardInfo: card.CardInfo{
					CardNumber: "8888888888",
					Pin: "888888",
					Balance: 900,
				},
				Name: "David Kurniawan",
				AccountNumber: "8888888888",
				BankBranch: "Semarang",
			},
		},
	}

	myAtm, err := cards.FindCardByAccountNumber("888888888")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myAtm.Name)
	}
	
	
	atm := card.ATM{
				CardInfo: card.CardInfo{
					CardNumber: "11111111111", 
					Pin: "111111", 
					Balance: 100000,
				},
				Name: "Dimas Saputro",
				AccountNumber: "1111111111",
				BankBranch: "Semarang",
			}

	atm.PrintDetails()

	st := transaction.SendTransaction{
		TransactionInfo: transaction.TransactionInfo{
			CardNumber: "2222222222",
			Date: "012323",
			Amount: 1000,
			Balance: 9000,
		},
		ReceiverAccountNumber: "1111111111",
		ReceiverName: "Dimas Saputro",
	}

	rt := transaction.ReceiveTransaction{
		TransactionInfo: transaction.TransactionInfo{
			CardNumber: "3333333333",
			Date: "012323",
			Amount: 20,
			Balance: 80,
		},
		SenderAccountNumber: "2222222222",
		SenderName: "Brian Rakajati",
	}

	transactions.Add(st)
	transactions.Add(rt)

	for v:= range transactions.Data {
		transactions.Data[v].PrintDetails()
	}

	filteredTransaction := transactions.FilterByCardNumber("2222222222")

	for i:= range filteredTransaction {
		filteredTransaction[i].PrintDetails()
	}

}