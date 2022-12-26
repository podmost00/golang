package bankApi

func ClientConstructor(name string, surname string, accountNumber string, cDeposit float64, cCredit float64) *Client {
	c := new(Client)

	c.name = name
	c.surname = surname
	c.accountNumber = accountNumber
	c.cDeposit = cDeposit
	c.cCredit = cCredit

	return c
}
