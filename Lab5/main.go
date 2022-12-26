package main

import (
	functions "Lab5_Structures_Interfaces/packages/productConsoleFunctions"
	"fmt"
)

func main() {
	fmt.Println("\nЗчитування продуктів...")
	products, err := functions.ReadProductsArray()

	if err != nil {
		fmt.Println("Під час зчитування масиву виникла помилка!")
	}

	fmt.Println("\nВиведення інформації про продукти...")
	functions.PrintProducts(products)

	fmt.Println("\nВизначення мінімального і максимального продуктів...")
	min, max := functions.GetProductsInfo(products)

	fmt.Println("Найдешевший продукт:", min.GetName(), fmt.Sprintf("%.2f", min.GetPriceIn()), "UAH")
	fmt.Println("Найдорожчий продукт:", max.GetName(), fmt.Sprintf("%.2f", max.GetPriceIn()), "UAH")
}
