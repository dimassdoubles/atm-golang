package machine

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"git.solusiteknologi.co.id/golab/atm/card"
	"git.solusiteknologi.co.id/golab/atm/transaction"
)

type Machine struct {
	CardRepo        *card.CardsRepository
	TransactionRepo *transaction.TransactionsRepository
	Card            card.Card
}

func (m *Machine) execute(menu int) error {
	var err error = nil
	switch menu {
	// print informasi kartu
	case 1:
		err = m.printInfo()
	// mutasi
	case 2:
		err = m.printMutasi()
	case 3:
		err = m.bayarBelanja()
	// tarik tunai
	case 4:
		err = m.tarik()
	// setor tunai
	case 5:
		err = m.setor()
	// transfer
	case 6:
		err = m.transfer()
	// topup e-money
	case 7:
		err = m.topupEmoney()
	default:
		fmt.Println("maaf, sistem tidak dapat memproses permintaan anda")
	}

	if err == nil {
		fmt.Println("transaksi berhasil")
	}

	fmt.Println()
	fmt.Println()

	return err
}

func (m Machine) Start() error {
	err := m.insert()
	if err != nil {
		return err
	}

	fmt.Println()
	menu := m.inputMenu()

	fmt.Println()
	err = m.execute(menu)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	for m.isNextTransaction() {
		nextMenu := m.inputMenu()

		err = m.execute(nextMenu)
		if err != nil {
			fmt.Println(err)
		}
	}

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
	var next string
	fmt.Print("ingin melakukan transaksi lagi ? (Y/n): ")
	fmt.Scanln(&next)

	if strings.TrimSpace(next) == "n" {
		m.Card = nil
		return false
	}

	// m.validate()

	return true
}

func (m Machine) setor() error {
	amount := inputAmount()

	err := m.Card.Add(amount)
	if err != nil {
		return err
	}

	transaction := transaction.SetorTransaction{
		Time:       getTime(),
		Nominal:    amount,
		CardNumber: m.Card.GetCardNumber(),
	}

	m.TransactionRepo.Add(transaction)

	return nil
}

func (m Machine) tarik() error {
	amount := inputAmount()

	err := m.Card.Substract(amount)
	if err != nil {
		return err
	}

	transaction := transaction.TarikTransaction{
		Time:       getTime(),
		Nominal:    amount,
		CardNumber: m.Card.GetCardNumber(),
	}

	m.TransactionRepo.Add(transaction)

	return nil
}

func (m Machine) printMutasi() error {

	transactions := m.TransactionRepo.FilterByCardNumber(m.Card.GetCardNumber())
	for i := range transactions {
		fmt.Println(transactions[i])
	}

	return nil
}

func (m Machine) printInfo() error {
	fmt.Println(m.Card)

	return nil
}

func (m Machine) bayarBelanja() error {
	amount := inputAmount()

	var err error = nil

	switch m.Card.(type) {
	case *card.ATM:
		if m.validate() {
			err = m.Card.Substract(amount)
			if err != nil {
				return err
			}
		}

	case *card.EMoney:
		err = m.Card.Substract(amount)
		if err != nil {
			return err
		}
	}

	transaction := transaction.BelanjaTransaction{
		Time:       getTime(),
		Nominal:    amount,
		CardNumber: m.Card.GetCardNumber(),
	}
	m.TransactionRepo.Add(transaction)

	return err
}

func (m Machine) transfer() error {
	accountNumber, err := inputCardNumber()
	if err != nil {
		return err
	}

	receiverCard, err := m.CardRepo.FindCardByAccountNumber(accountNumber)
	if err != nil {
		return err
	}

	amount := inputAmount()

	err = m.Card.Substract(amount)
	if err != nil {
		return err
	}
	err = receiverCard.Add(amount)
	if err != nil {
		m.Card.Add(amount)
		return err
	}

	sendTransaction := transaction.SendTransaction{
		Time:               getTime(),
		Nominal:            amount,
		CardNumber:         m.Card.GetCardNumber(),
		ReceiverCardNumber: receiverCard.GetCardNumber(),
	}

	receiveTransaction := transaction.ReceiveTransaction{
		Time:             getTime(),
		Nominal:          amount,
		CardNumber:       receiverCard.GetCardNumber(),
		SenderCardNumber: m.Card.GetCardNumber(),
	}

	m.TransactionRepo.Add(sendTransaction)
	m.TransactionRepo.Add(receiveTransaction)

	return nil
}

func (m Machine) topupEmoney() error {
	cardNumber, err := inputCardNumber()
	if err != nil {
		return err
	}

	receiverCard, err := m.CardRepo.FindCardByCardNumber(cardNumber)
	if err != nil {
		return err
	}

	receiverCard, isEmoney := receiverCard.(*card.EMoney)
	if !isEmoney {
		return errors.New("maaf, kartu e-money tidak terdaftar")
	}

	amount := inputAmount()

	err = m.Card.Substract(amount)
	if err != nil {
		return err
	}
	err = receiverCard.Add(amount)
	if err != nil {
		m.Card.Add(amount)
		return err
	}

	sendTransaction := transaction.SendTransaction{
		Time:               getTime(),
		Nominal:            amount,
		CardNumber:         m.Card.GetCardNumber(),
		ReceiverCardNumber: receiverCard.GetCardNumber(),
	}

	receiveTransaction := transaction.ReceiveTransaction{
		Time:             getTime(),
		Nominal:          amount,
		CardNumber:       receiverCard.GetCardNumber(),
		SenderCardNumber: m.Card.GetCardNumber(),
	}

	m.TransactionRepo.Add(sendTransaction)
	m.TransactionRepo.Add(receiveTransaction)

	return nil
}

func (m Machine) printMenu() {
	fmt.Println("pilihan menu: ")
	fmt.Println("1. informasi kartu")
	fmt.Println("2. mutasi")
	fmt.Println("3. pembayaran belanja")

	_, isAtm := m.Card.(*card.ATM)
	if isAtm {
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

	var menu int
	fmt.Print("masukan pilihan: ")
	_, err := fmt.Scan(&menu)

	if err != nil {
		fmt.Println(err)
		return m.inputMenu()
	}

	err = m.isMenuValid(menu)
	if err != nil {
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
	var cardNumber string
	fmt.Print("masukan card number: ")
	_, err := fmt.Scanln(&cardNumber)

	if err != nil {
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
	var pin string
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
	var amount int
	fmt.Print("masukan nominal: ")
	_, err := fmt.Scan(&amount)

	if err != nil {
		fmt.Println(err)
		return inputAmount()
	} else if amount <= 0 {
		fmt.Println("maaf, nominal tidak valid")
		return inputAmount()
	}

	return amount
}

func getTime() string {
	return time.Now().Format("2006-01-02 15:04")
}
