package card

import (
	"errors"
	"fmt"
)

type CardsRepository struct {
	Data []Card
}

func (cardRepo CardsRepository) FindCardByCardNumber(cardNumber string) (Card, error) {
	for i:= range cardRepo.Data {
		if cardRepo.Data[i].GetCardNumber() == cardNumber {
			return cardRepo.Data[i], nil
		}
	}
	return nil, errors.New("card number tidak terdaftar") 
}

func (cardRepo CardsRepository) FindCardByAccountNumber(accountNumber string) (*ATM, error) {
	fmt.Println("Masuk FindCardByAccountNumber")
	for i:= range cardRepo.Data {
		switch cardRepo.Data[i].(type) {
		case *ATM:
			if cardRepo.Data[i].(*ATM).AccountNumber == accountNumber {
				fmt.Println(cardRepo.Data[i].(*ATM).AccountNumber)
				fmt.Println(accountNumber)
				return cardRepo.Data[i].(*ATM), nil
			}
		}
	}
	return nil, errors.New("tidak ditemukan")
}