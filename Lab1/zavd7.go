package main

import "fmt"

func main() {

	//Задание.
	//1. Создайте 2 переменные  разных типов. Выпоните арифметические операции. Результат вывести
	variable8 := int8(10)
	variable16 := int16(10000)

	add := uint16(variable8) + uint16(variable16)
	fmt.Print("Операція додавання : ")
	fmt.Println(add)

	minus := uint16(variable16) - uint16(variable8)
	fmt.Print("Операція віднімання : ")
	fmt.Println(minus)

	multiply := uint32(variable16) * uint32(variable8)
	fmt.Print("Операція множення : ")
	fmt.Println(multiply)

	division := uint16(variable16) / uint16(variable8)
	fmt.Print("Операція ділення : ")
	fmt.Println(division)
}
