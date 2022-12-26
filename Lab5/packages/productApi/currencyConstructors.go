package productApi

func CurrencyConstructor(name string, exRate float64) *Currency {
	c := new(Currency)
	c.name = name
	c.exRate = exRate

	return c
}

func CurrencyConstructorWithExRate(exRate float64) *Currency {
	c := new(Currency)
	c.name = "USD"
	c.exRate = exRate

	return c
}

func CurrencyDefaultConstructor() *Currency {
	c := new(Currency)

	c.name = "USD"
	c.exRate = 36.84

	return c
}
