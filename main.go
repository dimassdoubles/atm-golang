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
		Data: []transaction.Transaction{},
	}

	cardRepo := &card.CardsRepository{
		Data: []card.Card{
			&card.ATM{
				CardInfo: card.CardInfo{
					CardNumber: "1111111111",
					Pin:        "111111",
					Balance:    85000,
				},
				Name:          "Dimas Saputro",
				AccountNumber: "1111111111",
				BankBranch:    "Semarang",
			},
			&card.ATM{
				CardInfo: card.CardInfo{
					CardNumber: "2222222222",
					Pin:        "222222",
					Balance:    280000,
				},
				Name:          "Siti Nurhaliza",
				AccountNumber: "2222222222",
				BankBranch:    "Semarang",
			},
			&card.ATM{
				CardInfo: card.CardInfo{
					CardNumber: "3333333333",
					Pin:        "333333",
					Balance:    10000,
				},
				Name:          "Budi Santoso",
				AccountNumber: "3333333333",
				BankBranch:    "Bandung",
			},
			&card.ATM{
				CardInfo: card.CardInfo{
					CardNumber: "4444444444",
					Pin:        "444444",
					Balance:    800000,
				},
				Name:          "Rina Anggraeni",
				AccountNumber: "4444444444",
				BankBranch:    "Medan",
			},
			&card.ATM{
				CardInfo: card.CardInfo{
					CardNumber: "5555555555",
					Pin:        "555555",
					Balance:    400000,
				},
				Name:          "Muhammad Yusuf",
				AccountNumber: "5555555555",
				BankBranch:    "Surabaya",
			},
			&card.EMoney{
				CardInfo: card.CardInfo{
					CardNumber: "9999999999",
					Pin:        "999999",
					Balance:    350000,
				},
			},
			&card.EMoney{
				CardInfo: card.CardInfo{
					CardNumber: "8888888888",
					Pin:        "888888",
					Balance:    650000,
				},
			},
		},
	}

	machine := machine.Machine{CardRepo: cardRepo, TransactionRepo: transactionRepo}

	for {
		err := machine.Start()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("transaksi selesai")
		fmt.Println()
		fmt.Println()
	}
}
