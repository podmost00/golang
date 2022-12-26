package strategy

import (
	"fmt"
	"math"
)

type BikeStrategy struct {
}

func (b *BikeStrategy) Route(distance float64) {
	speed := 20.2

	totalTime := distance * 1.2 / speed
	totalPrice := math.Round(distance*speed - totalTime)

	fmt.Printf("\nДоставка велосипедом на відстань %.3f триватиме %.2f год. та коштуватиме %.2f грн.", distance, totalTime, totalPrice)
}
