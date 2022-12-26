package productApi

import "fmt"

func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}

func (p *Product) SetCost(cost Currency) {
	p.cost = cost
}

func (p *Product) SetWeight(weight float64) {
	p.weight = weight
}

func (p *Product) SetProducer(producer string) {
	p.producer = producer
}

func (p *Product) SetQuantity(quantity int64) {
	p.quantity = quantity
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) GetCost() Currency {
	return p.cost
}

func (p *Product) GetQuantity() int64 {
	return p.quantity
}

func (p *Product) GetWeight() float64 {
	return p.weight
}

func (p *Product) GetProducer() string {
	return p.producer
}

func (p *Product) GetInfo() string {
	priceStr := fmt.Sprintf("%.2f", p.GetPrice())
	costInfo := p.cost.GetInfo()
	quantityStr := fmt.Sprintf("%d", p.GetQuantity())
	weightStr := fmt.Sprintf("%.2f", p.GetWeight())

	result := "Name: " + p.GetName() + "\n"
	result += "Price: " + priceStr + "\n"
	result += "Cost: " + costInfo + "\n"
	result += "Quantity: " + quantityStr + "\n"
	result += "Producer: " + p.GetProducer() + "\n"
	result += "Weight: " + weightStr + "\n"

	return result
}

func (p *Product) GetPriceIn() float64 {
	cost := p.GetCost()
	exRate := cost.GetExRate()

	return exRate * p.GetPrice()
}

func (p *Product) GetTotalPrice(isUah bool) float64 {
	if isUah {
		return p.GetPriceIn() * float64(p.GetQuantity())
	}

	return p.GetPrice() * float64(p.GetQuantity())
}

func (p *Product) GetTotalWeight() float64 {
	return p.GetWeight() * float64(p.GetQuantity())
}
