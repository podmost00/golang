package bankApi

func BankConstructor(name string, bankMoney float64, deposit float64, credit float64, clients []*Client) *Bank {
	b := new(Bank)

	b.name = name
	b.bankMoney = bankMoney
	b.deposit = deposit
	b.credit = credit
	b.clients = clients

	return b
}

func BankDefaultConstructor() *Bank {
	b := new(Bank)

	b.name = "Alpha"
	b.bankMoney = 56000000.0
	b.deposit = 0.0
	b.credit = 0.0
	b.clients = []*Client{}

	return b

}
