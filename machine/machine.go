package machine

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/golab/atm/card"
	"git.solusiteknologi.co.id/golab/atm/transaction"

	"regexp"
	"strings"
)

type Machine struct {
	CardRepo *card.CardsRepository
	TransactionRepo *transaction.TransactionsRepository
	Card card.Card
}

func (m *Machine) execute(menu int) error {
	var err error
	switch menu {
	// print informasi kartu
	case 1:
		m.printInfo()
		err = nil
	default:
		fmt.Println("Hello World!")
	}

	return err
}

func (m Machine) Start() error {
	err := m.insert()
	if err != nil {
		return err
	}

	menu := m.inputMenu()

	err = m.execute(menu)
	if err != nil {
		fmt.Println(err)
	}

	for m.isNextTransaction() {
		nextMenu := m.inputMenu()

		err = m.execute(nextMenu)
		if err != nil {
			return err
		}
	}

	fmt.Println("transaksi selesai")

	return nil
}	

func (m *Machine) insert() error {
	cardNumber, err := inputCardNumber()
	if err != nil {
		return err
	}

	selectedCard, err := m.CardRepo.FindCardByCardNumber(cardNumber)
	if err != nil {
		return err
	}

	pin := inputPin()

	if selectedCard.Validate(pin) {
		m.Card = selectedCard
		fmt.Println("selamat datang")
	} else {
		return errors.New("maaf, pin yang anda masukan salah")
	}
	
	return nil
}

func (m *Machine) isNextTransaction() bool {
	var next string;
	fmt.Print("ingin melakukan transaksi lagi ? (Y/n): ")
	_, err := fmt.Scanln(&next)
	
	if err != nil {
		return true
	}

	if strings.TrimSpace(next) == "n" {
		m.Card = nil
		return false
	}

	return true
}

func (m Machine) setor() error {

} 

func (m Machine) printInfo() {
	fmt.Println(m.Card)
}

func (m Machine) printMenu() {
	fmt.Println("pilihan menu: ")
	fmt.Println("1. informasi kartu")
	fmt.Println("2. mutasi")
	fmt.Println("3. pembayaran belanja")

	_, isAtm := m.Card.(*card.ATM)
	if (isAtm) {
		fmt.Println("4. tarik tunai")
		fmt.Println("5. setor tunai")
		fmt.Println("6. transfer")
		fmt.Println("7. topup e-money")
	}
}

func (m Machine) validate() bool {
	pin := inputPin()

	if !m.Card.Validate(pin) {
		fmt.Println("maaf, pin yang anda masukan salah")
		return m.validate()
	}
	return true
}

func (m Machine) inputMenu() int {
	m.printMenu()

	var menu int;
	fmt.Print("masukan pilihan: ")
	_, err := fmt.Scan(&menu)

	if err != nil {
		fmt.Println(err)
		return m.inputMenu()
	}

	err = m.isMenuValid(menu)
	if (err != nil) {
		fmt.Println(err)
		return m.inputMenu()
	}

	return menu
}

func (m Machine) isMenuValid(menu int) error {
	if menu < 1 {
		return errors.New("maaf, pilihan yang anda masukkan tidak valid")
	}

	switch m.Card.(type) {
	case *card.ATM:
		if menu > 7 {
			return errors.New("maaf, piihan yang anda masukan tidak valid")
		}
	case *card.EMoney:
		if menu > 3 {
			return errors.New("maaf, pilihan yang anda masukan tidak valid")
		}
	default:
		return nil
	}

	return nil
}



func isAtmValid(cardNumber string) bool {
	cardNumberLength := len(cardNumber)
	return cardNumberLength == 10 
}

func inputCardNumber() (string, error) {
	var cardNumber string;
	fmt.Print("masukan card number: ")
	_, err := fmt.Scanln(&cardNumber)

	if (err != nil) {
		return "", err
	} else if isAtmValid(cardNumber) {
		return cardNumber, nil
	}
	return "", errors.New("card number tidak valid")
}

func isPinValid(pin string) error {
	re := regexp.MustCompile(`^\d+$`)
	if re.MatchString(pin) {
		return nil
	}
	return errors.New("maaf, pin tidak valid")
}

func inputPin() string {
	var pin string;
	fmt.Print("masukan pin: ")
	_, err := fmt.Scanln(&pin)

	if err != nil {
		fmt.Println(err)
		return inputPin()
	}

	err = isPinValid(pin)

	if err != nil {
		fmt.Println(err)
		return inputPin()
	}

	return pin
}

func inputAmount() int {
	var amount int;
	fmt.Print("masukan nominal: ")
	_, err := fmt.Scan(&amount)

	if err != nil {
		fmt.Println(err)
		return inputAmount()
	}

	return amount
}
