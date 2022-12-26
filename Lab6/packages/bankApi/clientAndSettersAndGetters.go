package bankApi

func (c *Client) SetName(name string) {
	c.name = name
}

func (c *Client) GetName() string {
	return c.name
}

func (c *Client) SetSurname(surname string) {
	c.surname = surname
}

func (c *Client) GetSurname() string {
	return c.surname
}

func (c *Client) GetFullName() string {
	return c.surname + " " + c.name
}

func (c *Client) SetAccountNumber(accountNumber string) {
	c.accountNumber = accountNumber
}

func (c *Client) GetAccountNumber() string {
	return c.accountNumber
}

func (c *Client) SetCDeposit(cDeposit float64) {
	c.cDeposit = cDeposit
}

func (c *Client) GetCDeposit() float64 {
	return c.cDeposit
}

func (c *Client) SetCCredit(cCredit float64) {
	c.cCredit = cCredit
}

func (c *Client) GetCCredit() float64 {
	return c.cCredit
}

func (c *Client) GetBank() *Bank {
	return c.bank
}

func (c *Client) SetBank(bank *Bank) {
	c.bank = bank
}
