package card

import (
	"errors"
)

type CardInfo struct {
	CardNumber, Pin string
	Balance float64
}

type Card interface {
	Add(amount float64) error
	Substract(amount float64) error
	PrintDetails()
	GetCardNumber() string
}

type ATM struct {
	CardInfo
	Name, AccountNumber, BankBranch string
}

func (atm *ATM) Add(amount float64) error {
	atm.Balance = atm.Balance + amount
	return nil
}

func (atm *ATM) Substract(amount float64) error {
	if amount < atm.Balance {
		atm.Balance = atm.Balance - amount
		return nil
	}
	return errors.New("saldo tidak cukup")
}

func (atm ATM) GetCardNumber() string {
	return atm.CardNumber
}

type EMoney struct {
	CardInfo
}

func (emoney *EMoney) Add(amount float64) error {
	newBalance := emoney.CardInfo.Balance + amount
	if newBalance > 1000000 {
		return errors.New("saldo melebihi batas maksimal")
	}
	emoney.Balance = newBalance
	return nil
}

func (emoney *EMoney) Substract(amount float64) error {
	if amount < emoney.Balance {
		emoney.Balance = emoney.Balance - amount
		return nil
	}
	return errors.New("saldo tidak cukup")
}

func (emoney EMoney) GetCardNumber() string {
	return emoney.CardNumber
}
