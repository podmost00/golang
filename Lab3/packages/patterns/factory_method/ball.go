package factory_method

import "fmt"

type Ball struct {
	Type  string
	Size  int
	Color string
}

func CreateBall() Category {
	return Ball{
		Type:  BallType,
		Size:  5,
		Color: "white-black",
	}
}

func (b Ball) GetType() string {
	return b.Type
}

func (b Ball) GetInfo() string {
	return fmt.Sprintf("Type: %s Size: %d Color: %s", b.Type, b.Size, b.Color)
}
