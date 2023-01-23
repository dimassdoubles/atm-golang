package repo

import (
	"errors"
	"git.solusiteknologi.co.id/golab/atm/card"
)

type CardsRepository struct {
	Data []card.Card
}

func (cr CardsRepository) FindCardByCardNumber(cardNumber string) card.Card {
	for i:= range cr.Data {
		if cr.Data[i].GetCardNumber() == cardNumber {
			return cr.Data[i]
		}
	}
	return nil
}

func (cr CardsRepository) FindCardByAccountNumber(accountNumber string) (*card.ATM, error) {
	for i:= range cr.Data {
		switch cr.Data[i].(type) {
		case *card.ATM:
			if cr.Data[i].(*card.ATM).AccountNumber == accountNumber {
				return cr.Data[i].(*card.ATM), nil
			}
		}
	}
	return nil, errors.New("tidak ditemukan")
}