package bankApi

import "fmt"

func CreateNewClient(name string, surname string, accountNumber string, deposit float64, credits float64) Client {
	return Client{name, surname, accountNumber, deposit, credits, nil}
}

func (c *Client) PrintInfo() {
	fmt.Printf("\nCLIENT INFO\nFullName: %-10s\nDeposit: %-10.2f\nCredit:%-10.2f\nAccount number:%-10s\n", c.GetFullName(), c.GetCDeposit(), c.GetCCredit(), c.GetAccountNumber())
}

func (c *Client) CheckMoney(m float64) bool {
	if c.GetCDeposit() >= m {
		return true
	}

	fmt.Printf("%s, у вас недостатньо коштів: %.3f\n", c.GetFullName(), m)
	fmt.Printf("Ви маєте: %.3f\n", c.GetCDeposit())
	return false
}

func (c *Client) AddToDeposit(deposit float64) {
	c.cDeposit += deposit

	c.bank.AddToDeposit(deposit)
}

func (c *Client) SubFromDeposit(deposit float64) {
	c.cDeposit -= deposit

	c.bank.SubFromDeposit(deposit)
}

func (c *Client) AddToCredit(credit float64) {
	c.cCredit -= credit

	c.bank.AddToCredit(credit)
}

func (c *Client) PayCredit(money float64) {
	if money < 0 {
		return
	}

	if !c.CheckMoney(money) {
		return
	}

	if money > -c.GetCCredit() {
		extra := money + c.GetCCredit()
		money -= extra
	}

	c.SubFromDeposit(money)
	c.SubFromCredit(money)
}

func (c *Client) TakeCredit(money float64) {
	if !c.bank.CheckMoney(money) {
		return
	}

	c.AddToDeposit(money)
	c.AddToCredit(money)
}

func (c *Client) SubFromCredit(credit float64) {
	c.cCredit += credit

	c.bank.SubFromCredit(credit)
}

func (c *Client) TakeFromDeposit(money float64) {
	if !c.CheckMoney(money) {
		return
	}

	c.SubFromDeposit(money)
}
