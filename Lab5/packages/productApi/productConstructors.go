package productApi

func ProductConstructor(name string, price float64, cost Currency,
	quantity int64, producer string, weight float64) *Product {
	p := new(Product)

	p.name = name
	p.price = price
	p.cost = cost
	p.quantity = quantity
	p.producer = producer
	p.weight = weight

	return p
}

func ProductConstructorWithoutCost(name string, price float64,
	quantity int64, producer string, weight float64) *Product {
	p := new(Product)

	p.name = name
	p.price = price
	p.cost = Currency{"USD", 36.8}
	p.quantity = quantity
	p.producer = producer
	p.weight = weight

	return p
}

func ProductDefaultConstructor() *Product {
	p := new(Product)

	p.name = "Tab M10 (3rd Gen) 4/64 Wi-Fi Storm Grey"
	p.price = 230.7
	p.cost = Currency{"USD", 36.8}
	p.quantity = 54
	p.producer = "Lenovo"
	p.weight = 423.7

	return p
}
