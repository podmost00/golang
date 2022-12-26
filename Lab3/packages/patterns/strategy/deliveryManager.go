package strategy

type DeliveryManager struct {
	Delivery
}

func (m *DeliveryManager) SetDelivery(delivery Delivery) {
	m.Delivery = delivery
}
