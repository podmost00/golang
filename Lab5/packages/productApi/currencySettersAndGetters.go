package productApi

import "fmt"

func (c *Currency) SetExRate(exRate float64) {
	c.exRate = exRate
}

func (c *Currency) SetName(name string) {
	c.name = name
}

func (c *Currency) GetExRate() float64 {
	return c.exRate
}

func (c *Currency) GetName() string {
	return c.name
}

func (c *Currency) GetInfo() string {
	exRateStr := fmt.Sprint(c.exRate)

	return "1 " + c.name + " = " + exRateStr + " UAH"
}
