package transaction

import (
	"fmt"
)

type Transaction interface {
	GetCardNumber() string
}

type SetorTransaction struct {
	Time string
	Nominal float64
	CardNumber string
}

func (t SetorTransaction) GetCardNumber() string {
	return t.CardNumber
}

func (t SetorTransaction) String() string {
	return fmt.Sprintln(t.Time, " | ", t.Nominal, " - setor tunai");
}

type TarikTransaction struct {
	Time string
	Nominal float64
	CardNumber string
}

func (t TarikTransaction) GetCardNumber() string {
	return t.CardNumber
}

func (t TarikTransaction) String() string {
	return fmt.Sprintln(t.Time, " | ", t.Nominal, " - tarik tunai")
}

type BelanjaTransaction struct {
	Time string
	Nominal float64
	CardNumber string
}

func (t BelanjaTransaction) GetCardNumber() string {
	return t.CardNumber
}

func (t BelanjaTransaction) String() string {
	return fmt.Sprintln(t.Time, " | ", t.Nominal, " - pembayaran belanja")
}

type SendTransaction struct {
	Time string
	Nominal float64
	CardNumber string
	ReceiverCardNumber string
}

func (t SendTransaction) GetCardNumber() string {
	return t.CardNumber
}

func (t SendTransaction) String() string {
	return fmt.Sprintln(t.Time, " | ", t.Nominal, " - kirim ke ", t.ReceiverCardNumber)
}

type ReceiveTransaction struct {
	Time string
	Nominal float64
	CardNumber string
	SenderCardNumber string
}

func (t ReceiveTransaction) GetCardNumber() string {
	return t.CardNumber
}

func (t ReceiveTransaction) String() string {
	return fmt.Sprintln(t.Time, " | ", t.Nominal, " - terima dari ", t.SenderCardNumber)
}

type TopUpEmoneyTransaction struct {
	Time string
	Nominal float64
	CardNumber string
	ReceiverCardNumber string
}

func (t TopUpEmoneyTransaction) GetCardNumber() string {
	return t.CardNumber
}

func (t TopUpEmoneyTransaction) String() string {
	return fmt.Sprintln(t.Time, " | ", t.Nominal, " - top up e-money ", t.ReceiverCardNumber)
}

