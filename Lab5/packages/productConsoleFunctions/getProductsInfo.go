package productConsoleFunctions

import "Lab5_Structures_Interfaces/packages/productApi"

func GetProductsInfo(products []productApi.Product) (min productApi.Product, max productApi.Product) {
	min = products[0]
	max = products[0]

	for i := 1; i < len(products); i++ {
		if products[i].GetPriceIn() < min.GetPriceIn() {
			min = products[i]
		} else if products[i].GetPriceIn() > max.GetPriceIn() {
			max = products[i]
		}
	}

	return min, max
}
