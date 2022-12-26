package factory_method

import "errors"

/*
	У нашому маркетплейсі футбольної екіпіровки є декілька видів товарів: м'ячі, воротарські рукавиці та футболки.
	Завдяки фабричному методу у нас буде єдиний інтерфейс об'єктів різних категорій.
	Це зручно, адже з часом кількість категорій може збільшуватись
*/

const (
	BallType            = "ball"
	GoalkeeperEquipType = "goalkeeper"
	TShortsType         = "t-short"
)

type Category interface {
	GetType() string
	GetInfo() string
}

func Create(typeName string) (Category, error) {
	switch typeName {
	default:
		return nil, errors.New("category not found")
	case BallType:
		return CreateBall(), nil
	case GoalkeeperEquipType:
		return CreateGoalkeeperEquip(), nil
	case TShortsType:
		return CreateTShort(), nil
	}
}
