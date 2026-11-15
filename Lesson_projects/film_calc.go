package main

import "fmt"

func main() {
	fmt.Printf("Цена если только ночью:%d\n", ticetPrice(100, true, false))
	fmt.Printf("Цена если ночью в выходные:%d\n", ticetPrice(100, true, true))
	fmt.Printf("Цена если днем в выходные:%d\n", ticetPrice(100, false, true))
	fmt.Printf("Цена если днем в будни:%d\n", ticetPrice(100, false, false))
}

func ticetPrice(basePrice int, nightPrice bool, holydaysPrice bool) int {
	finalPrice := basePrice

	if nightPrice && holydaysPrice {
		finalPrice += 150
	} else if nightPrice {
		finalPrice += 50
	} else if holydaysPrice {
		finalPrice += 100
	} else {
		finalPrice = basePrice
	}

	return finalPrice
}
