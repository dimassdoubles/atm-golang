package main

import (
	// "errors"
	"fmt"

	"git.solusiteknologi.co.id/golab/atm/card"
	"git.solusiteknologi.co.id/golab/atm/machine"
	"git.solusiteknologi.co.id/golab/atm/transaction"
)

// func isAtmValid(accountNumber string) bool {
// 	accountnumberLength := len(accountNumber)
// 	return accountnumberLength == 10
// }

// func inputAccountNumber() (string, error) {
// 	var accountNumber string;
// 	fmt.Print("Masukan account number: ")
// 	_, err := fmt.Scanln(&accountNumber)

// 	if (err != nil) {
// 		return "", err
// 	} else if isAtmValid(accountNumber) {
// 		return accountNumber, nil
// 	}
// 	return "", errors.New("account number tidak valid")
// }

func main() {

	transactionRepo := &transaction.TransactionsRepository{
		Data : []transaction.Transaction{},
	}

	cardRepo := &card.CardsRepository{
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
			&card.ATM{
				CardInfo: card.CardInfo{
					CardNumber: "1111111111",
					Pin: "111111",
					Balance: 1000,
				},
				Name: "David Kurniawan",
				AccountNumber: "1111111111",
				BankBranch: "Semarang",
			},
		},
	}

	machine := machine.Machine{CardRepo: cardRepo, TransactionRepo: transactionRepo}

	for {
		err := machine.Start()
		if (err != nil) {
			fmt.Println(err)
		} else {
			fmt.Println("tidak ada error")
		}
	}
	


	// fmt.Println(cards.Data[0].GetCardNumber());

	// accountNumber, err := inputAccountNumber()
	// if err != nil {
	// 	panic(err)
	// }
	
	// myAtm, err := cards.FindCardByAccountNumber(accountNumber)
	// if err != nil {
	// 	panic(err)
	// }

	// myAtm.PrintDetails()

	// st := transaction.SendTransaction{
	// 	TransactionInfo: transaction.TransactionInfo{
	// 		CardNumber: "2222222222",
	// 		Date: "012323",
	// 		Amount: 1000,
	// 		Balance: 9000,
	// 	},
	// 	ReceiverAccountNumber: "1111111111",
	// 	ReceiverName: "Dimas Saputro",
	// }

	// rt := transaction.ReceiveTransaction{
	// 	TransactionInfo: transaction.TransactionInfo{
	// 		CardNumber: "3333333333",
	// 		Date: "012323",
	// 		Amount: 20,
	// 		Balance: 80,
	// 	},
	// 	SenderAccountNumber: "2222222222",
	// 	SenderName: "Brian Rakajati",
	// }

	// transactions.Add(st)
	// transactions.Add(rt)

	// for v:= range transactions.Data {
	// 	transactions.Data[v].PrintDetails()
	// }

	// filteredTransaction := transactions.FilterByCardNumber("2222222222")

	// for i:= range filteredTransaction {
	// 	filteredTransaction[i].PrintDetails()
	// }

}