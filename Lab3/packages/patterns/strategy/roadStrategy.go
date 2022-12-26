package strategy

import (
	"fmt"
	"math"
)

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(distance float64) {
	speed := 65.6
	traffic := 1.7

	totalTime := distance * traffic / speed
	totalPrice := math.Round(distance*speed - totalTime + traffic)

	fmt.Printf("\nДоставка автомобілем на відстань %.3f триватиме %.2f год. та коштуватиме %.2f грн.", distance, totalTime, totalPrice)
}
