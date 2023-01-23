package transaction

type Transaction interface {
	PrintDetails()
	GetCardNumber() string
}

type TransactionInfo struct {
	CardNumber, Date string
	Amount, Balance float64
}

type SendTransaction struct {
	TransactionInfo
	ReceiverAccountNumber, ReceiverName string
}

func (st SendTransaction) GetCardNumber() string {
	return st.CardNumber
}

type ReceiveTransaction struct {
	TransactionInfo
	SenderAccountNumber, SenderName string
}

func (rt ReceiveTransaction) GetCardNumber() string {
	return rt.CardNumber
}

type TopUpTransaction struct {
	TransactionInfo
	ReceiverCardNumber string
}

func (tt TopUpTransaction) GetCardNumber() string {
	return tt.CardNumber
}

type CheckOutTransaction struct {
	TransactionInfo
}

func (ct CheckOutTransaction) GetCardNumber() string {
	return ct.CardNumber
}

