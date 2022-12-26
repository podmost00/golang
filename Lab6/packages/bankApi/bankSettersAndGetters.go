package bankApi

func (b *Bank) SetName(name string) {
	b.name = name
}

func (b *Bank) GetName() string {
	return b.name
}

func (b *Bank) SetBankMoney(bankMoney float64) {
	b.bankMoney = bankMoney
}

func (b *Bank) GetBankMoney() float64 {
	return b.bankMoney
}

func (b *Bank) SetDeposit(deposit float64) {
	b.deposit = deposit
}

func (b *Bank) GetDeposit() float64 {
	return b.deposit
}

func (b *Bank) SetCredit(credit float64) {
	b.credit = credit
}

func (b *Bank) GetCredit() float64 {
	return b.credit
}

func (b *Bank) SetClients(clients []*Client) {
	b.clients = clients
}

func (b *Bank) GetClients() []*Client {
	return b.clients
}
