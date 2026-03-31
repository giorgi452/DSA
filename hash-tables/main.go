package main

import "fmt"

func main() {
	products := map[string]float32{
		"Milk":    1.49,
		"Avocado": 1.49,
		"Apple":   0.67,
	}

	fmt.Println("Product Prices: ", products)
	fmt.Println("Avocado Price: ", products["Avocado"])
}
