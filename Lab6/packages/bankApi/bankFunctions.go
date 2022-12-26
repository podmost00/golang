package bankApi

import "fmt"

func CreateNewBank(name string, bankMoney float64) Bank {
	return Bank{name, bankMoney, 0.0, 0.0, []*Client{}}
}

func (b *Bank) TakeMoney(money float64) {
	b.bankMoney -= money
}

func (b *Bank) PutMoney(money float64) {
	b.bankMoney -= money
}

func (b *Bank) CheckMoney(m float64) bool {
	if b.GetBankMoney() >= m {
		return true
	}

	return false
}

func (b *Bank) AddToDeposit(deposit float64) {
	b.deposit += deposit
}

func (b *Bank) SubFromDeposit(deposit float64) {
	b.deposit -= deposit
}

func (b *Bank) AddToCredit(credit float64) {
	b.credit += credit

	b.TakeMoney(credit)
}

func (b *Bank) SubFromCredit(credit float64) {
	b.credit -= credit

	b.PutMoney(credit)
}

func (b *Bank) GetClient(index int) *Client {
	return b.clients[index]
}

func (b *Bank) PrintClientByAccountNumber(acc string) {
	for _, c := range b.clients {
		if c.accountNumber == acc {
			c.PrintInfo()
			break
		}
	}
}
func (b *Bank) PrintClientsBySurname(surname string) {
	for _, c := range b.clients {
		if c.surname == surname {
			c.PrintInfo()
		}
	}
}

func (b *Bank) AddClient(c *Client) {
	b.clients = append(b.clients, c)

	b.AddToDeposit(c.cDeposit)
	b.AddToCredit(-c.cCredit)

	c.bank = b
}

func (b *Bank) RemoveClient(id int) {
	b.clients = append(b.clients[:id], b.clients[id+1:]...)
}

func (b *Bank) PrintInfo() {
	fmt.Printf("Банк: %s\nДепозити: %.2f\nКредити: %.2f\nБаланс: %.2f", b.name, b.deposit, b.credit, b.bankMoney)

	fmt.Println("\nКлієнти")
	for _, c := range b.clients {
		c.PrintInfo()
	}
}
