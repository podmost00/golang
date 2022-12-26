package bankApi

type Client struct {
	name          string
	surname       string
	accountNumber string
	cDeposit      float64
	cCredit       float64
	bank          *Bank
}
