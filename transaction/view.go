package transaction
import(
	"fmt"
)

func printDetails(ti TransactionInfo, description string) {
	fmt.Println("- ", description, " - Nominal: ", ti.Amount, " - Saldo: ", ti.Balance)
}

func (st SendTransaction) PrintDetails() {
	description := fmt.Sprintf("Mengirim transfer ke %v %v", st.ReceiverAccountNumber, st.ReceiverName)
	printDetails(st.TransactionInfo, description)
}

func (rt ReceiveTransaction) PrintDetails() {
	description := fmt.Sprintf("Menerima transfer dari %v %v", rt.SenderAccountNumber, rt.SenderName)
	printDetails(rt.TransactionInfo, description)
}