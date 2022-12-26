package main

//Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	//Задание.
	//1. Создайте переменные разных типов, используя краткую запись и инициализацию по-умолчанию. Результат вывести
	exampleInt := 10
	fmt.Println(exampleInt)
	exampleFloat := -3.4
	fmt.Println(exampleFloat)
	exampleBoolean := true
	fmt.Println(exampleBoolean)
	exampleSimbol := 'A'
	fmt.Println(exampleSimbol)
	exampleString := "GoLang"
	fmt.Println(exampleString)
}
