package strategy

import (
	"fmt"
	"math"
)

type CourierStrategy struct {
}

func (b *CourierStrategy) Route(distance float64) {
	speed := 3.2

	totalTime := distance / speed
	totalPrice := math.Round(distance / speed * totalTime)

	fmt.Printf("\nДоставка кур'єром на відстань %.3f триватиме %.2f год. та коштуватиме %.2f грн.", distance, totalTime, totalPrice)
}
