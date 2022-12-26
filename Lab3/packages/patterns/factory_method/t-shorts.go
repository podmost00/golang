package factory_method

import "fmt"

type TShort struct {
	Type   string
	Size   int
	Color  string
	Gender string
}

func CreateTShort() Category {
	return TShort{
		Type:   TShortsType,
		Size:   5,
		Color:  "white-black",
		Gender: "male",
	}
}

func (t TShort) GetType() string {
	return t.Type
}

func (t TShort) GetInfo() string {
	return fmt.Sprintf("Type: %s Size: %d Color: %s Gender: %s", t.Type, t.Size, t.Color, t.Gender)
}
