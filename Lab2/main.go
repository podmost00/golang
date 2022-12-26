package main

import (
	"Lab2_Packages_Testing/packages/math"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Лаб_2. Оберіть операцію: \n\t1) Знаходження мінімального значення з трьох елементів;\n\t2) Обчисленнясереднього значення трьох елементів ;\n\t3) Рішення рівняння першого порядку.")
	var menu int
	fmt.Print("\nВаш вибір: ")
	fmt.Fscan(os.Stdin, &menu)
	switch menu {
	case 1:
		fmt.Println("Обрано знаходження мінімального елементу.")
		slice := readSlice()
		result, err := math.FindMin(slice...)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\nМінімальний елемент:", result)
		}
	case 2:
		fmt.Println("Обрано знаходження середнього арифметичного.")
		slice := readSlice()
		result, err := math.GetAverage(slice...)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\nСереднє арифметичне:", result)
		}
	case 3:
		fmt.Println("Обрано вирішення лінійного рівняння.")
		var k, x, b float64
		fmt.Print("\nВведіть k: ")
		fmt.Fscan(os.Stdin, &k)
		fmt.Print("Введіть x: ")
		fmt.Fscan(os.Stdin, &x)
		fmt.Print("Введіть b: ")
		fmt.Fscan(os.Stdin, &b)
		result := math.SolveLinearEquation(k, x, b)
		fmt.Println("Результат вирішення лінійного рівняння:", result)
	default:
		fmt.Println("Такого пункту меню не існує")
	}
}

func readSlice() []float64 {
	var size int
	fmt.Print("\nВведіть розмірність: ")
	fmt.Fscan(os.Stdin, &size)
	if size == 0 {
		fmt.Println("Введено некоректну розмірність.")
		return nil
	}
	slice := make([]float64, size)
	fmt.Println("\nПорядково введіть елементи: ")
	for i := 0; i < size; i++ {
		var number float64
		fmt.Fscan(os.Stdin, &number)
		slice[i] = number
	}
	return slice
}
