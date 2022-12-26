package bankApi

import (
	"fmt"
	"math/rand"
	"time"
)

func Deposit(c *Client, ch chan struct{}) {
	for c.GetCDeposit() > 0 {
		choice := rand.Intn(2) + 1

		<-ch

		sum := 200.0
		switch choice {
		case 1:
			{
				c.AddToDeposit(sum)
				break
			}
		case 2:
			{
				c.TakeFromDeposit(sum)
				break
			}
		}

		ch <- struct{}{}

		time.Sleep(time.Second)
	}

	fmt.Println("\n" + c.GetFullName() + ", роботу з депозитами завершено\n")
}

func Credit(c *Client, ch chan struct{}) {
	for c.GetCCredit() < 0 {
		choice := rand.Intn(2) + 1

		<-ch

		sum := 200.0
		switch choice {
		case 1:
			{
				c.TakeCredit(sum)
				break
			}
		case 2:
			{
				c.PayCredit(sum)
				break
			}
		}

		ch <- struct{}{}

		time.Sleep(time.Second)
	}

	fmt.Println("\n" + c.GetFullName() + ", роботу з кредитами завершено\n")
}
