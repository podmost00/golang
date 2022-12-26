package productConsoleFunctions

import (
	"Lab5_Structures_Interfaces/packages/productApi"
	"errors"
	"fmt"
	"os"
)

func ReadProductsArray() ([]productApi.Product, error) {
	var products []productApi.Product

	var size int
	fmt.Print("Введіть кількість елементів у масиві продуктів: ")
	_, err := fmt.Fscan(os.Stdin, &size)

	if err != nil {
		return nil, errors.New("incorrect array length value")
	}

	for i := 0; i < size; i++ {
		fmt.Println("\nЗаповнення інформації про продукт #", i+1)

		item := readProduct()
		products = append(products, *item)
	}

	return products, nil
}

func readProduct() *productApi.Product {
	fmt.Println("Оберіть тип конструктора структури продукту: \n1) З усіма полями \n2) Валюта за замочуванням \n3) За замовчуванням")

	var constructorType int64
	fmt.Print("Введіть обраний варіант: ")
	_, err := fmt.Fscan(os.Stdin, &constructorType)

	var name string
	var price float64
	var cost *productApi.Currency
	var quantity int64
	var producer string
	var weight float64

	switch constructorType {
	case 1:
		fmt.Print("Введіть назву продукту: ")
		_, err = fmt.Fscan(os.Stdin, &name)

		fmt.Print("Введіть ціну продукту: ")
		_, err = fmt.Fscan(os.Stdin, &price)

		cost = readCurrency()

		fmt.Print("Введіть кількість продукту: ")
		_, err = fmt.Fscan(os.Stdin, &quantity)

		fmt.Print("Введіть назву виробника продукту: ")
		_, err = fmt.Fscan(os.Stdin, &producer)

		fmt.Print("Введіть вагу продукту: ")
		_, err = fmt.Fscan(os.Stdin, &weight)

		if err != nil {
			fmt.Println("Помилка під час введення! Обрано конструктор за замовчуванням")
			return productApi.ProductDefaultConstructor()
		}

		return productApi.ProductConstructor(name, price, *cost, quantity, producer, weight)
	case 2:
		fmt.Print("Введіть назву продукту: ")
		_, err = fmt.Fscan(os.Stdin, &name)

		fmt.Print("Введіть ціну продукту: ")
		_, err = fmt.Fscan(os.Stdin, &price)

		fmt.Print("Введіть кількість продукту: ")
		_, err = fmt.Fscan(os.Stdin, &quantity)

		fmt.Print("Введіть назву виробника продукту: ")
		_, err = fmt.Fscan(os.Stdin, &producer)

		fmt.Print("Введіть вагу продукту: ")
		_, err = fmt.Fscan(os.Stdin, &weight)

		if err != nil {
			fmt.Println("Помилка під час введення! Обрано конструктор за замовчуванням")
			return productApi.ProductDefaultConstructor()
		}

		return productApi.ProductConstructorWithoutCost(name, price, quantity, producer, weight)
	case 3:
		return productApi.ProductDefaultConstructor()
	default:
		fmt.Println("Помилка! Неправильно обраний варіант, обрано конструктор за замовчуванням")
		return productApi.ProductDefaultConstructor()
	}
}

func readCurrency() *productApi.Currency {
	fmt.Println("Оберіть тип конструктора структури валюти: \n1) З назвою та курсом \n2) Лише курсом \n3) За замовчуванням")

	var constructorType int64
	fmt.Print("Введіть обраний варіант: ")
	_, err := fmt.Fscan(os.Stdin, &constructorType)

	var name string
	var exRate float64

	switch constructorType {
	case 1:
		fmt.Print("Введіть назву валюти: ")
		_, err = fmt.Fscan(os.Stdin, &name)

		fmt.Print("Введіть курс: ")
		_, err = fmt.Fscan(os.Stdin, &exRate)

		if err != nil {
			fmt.Println("Помилка під час введення! Обрано конструктор за замовчуванням")
			return productApi.CurrencyDefaultConstructor()
		}

		return productApi.CurrencyConstructor(name, exRate)
	case 2:
		fmt.Print("Введіть курс: ")
		_, err = fmt.Fscan(os.Stdin, &exRate)

		if err != nil {
			fmt.Println("Помилка під час введення! Обрано конструктор за замовчуванням")
			return productApi.CurrencyDefaultConstructor()
		}

		return productApi.CurrencyConstructorWithExRate(exRate)
	case 3:
		return productApi.CurrencyDefaultConstructor()
	default:
		fmt.Println("Помилка! Неправильно обраний варіант, обрано конструктор за замовчуванням")
		return productApi.CurrencyDefaultConstructor()
	}
}
