package card

import (
	"errors"
	"fmt"
)

type CardInfo struct {
	CardNumber, Pin string
	Balance float64
}

type Card interface {
	Add(amount float64) error
	Substract(amount float64) error
	GetCardNumber() string
	Validate(pin string) bool
}

type ATM struct {
	CardInfo
	Name, AccountNumber, BankBranch string
}

func (c CardInfo) String() string {
	result := fmt.Sprintln("Informasi Kartu")
	result += fmt.Sprintln("---------------")
	result += fmt.Sprintln("Nomor Kartu    :", c.CardNumber)
	result += fmt.Sprintln("Saldo          :", c.Balance)

	return result
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

func (atm *ATM) Validate(pin string) bool {
	return pin == atm.Pin
}

func (atm *ATM) GetCardNumber() string {
	return atm.CardNumber
}

func (atm *ATM) String() string {
	result := atm.CardInfo.String()
	result += fmt.Sprintln("Nomor Rekening :", atm.AccountNumber)
	result += fmt.Sprintln("Cabang Bank    :", atm.BankBranch)
	result += fmt.Sprintln("Nama Pemilik   :", atm.Name)
	return result
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

func (emoney *EMoney) GetCardNumber() string {
	return emoney.CardNumber
}

func (emoney *EMoney) Validate(pin string) bool {
	return pin == emoney.Pin
}

func (emoney *EMoney) String() string {
	return emoney.CardInfo.String()
}