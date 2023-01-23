package card

import (
	"fmt"
)

func (ci CardInfo) printBaseInfo(cardType string) {
	fmt.Println("Informasi Kartu")
	fmt.Println("---------------")
	fmt.Println("Tipe Kartu     :", cardType)
	fmt.Println("Nomor Kartu    :", ci.CardNumber)
	fmt.Println("Saldo          :", ci.Balance)
}

func (atm ATM) PrintDetails() {
	atm.printBaseInfo("ATM")
	fmt.Println("Nomor Rekening :", atm.AccountNumber)
	fmt.Println("Cabang Bank    :", atm.BankBranch)
	fmt.Println("Nama Pemilik   :", atm.Name)
}

func (emoney EMoney) PrintDetails() {
	emoney.printBaseInfo("E-Money")
}
