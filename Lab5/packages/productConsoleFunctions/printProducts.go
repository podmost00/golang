package productConsoleFunctions

import (
	"Lab5_Structures_Interfaces/packages/productApi"
	"fmt"
)

func PrintProducts(products []productApi.Product) {
	for index, item := range products {
		fmt.Println("\nПродукт #", index+1)
		PrintProduct(item)
	}
}

func PrintProduct(p productApi.Product) {
	fmt.Println(p.GetInfo())
}
