package bankApi

type Bank struct {
	name      string
	bankMoney float64
	deposit   float64
	credit    float64
	clients   []*Client
}
