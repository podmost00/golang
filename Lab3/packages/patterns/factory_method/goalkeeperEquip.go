package factory_method

import "fmt"

type GoalkeeperEquip struct {
	Type  string
	Size  int
	Brand string
}

func CreateGoalkeeperEquip() Category {
	return GoalkeeperEquip{
		Type:  GoalkeeperEquipType,
		Size:  8,
		Brand: "Nike",
	}
}

func (g GoalkeeperEquip) GetType() string {
	return g.Type
}

func (g GoalkeeperEquip) GetInfo() string {
	return fmt.Sprintf("Type: %s Size: %d Brand: %s", g.Type, g.Size, g.Brand)
}