package main

import "fmt"

func main() {
	fmt.Printf("Цена если только ночью:%d\n", ticetPrice(100, true, false))
	fmt.Printf("Цена если ночью в выходные:%d\n", ticetPrice(100, true, true))
	fmt.Printf("Цена если днем в выходные:%d\n", ticetPrice(100, false, true))
	fmt.Printf("Цена если днем в будни:%d\n", ticetPrice(100, false, false))
	receipt, total := formatReceipt("IT", 10, 100) //Распаковка значений receipt, total
	fmt.Printf("Ваш билет\n%sПолная стоимость: %d", receipt, total)
}

func ticetPrice(basePrice int, isEvening bool, isWeekend bool) int {
	price := basePrice

	if isEvening {
		price += 50
	}
	if isWeekend {
		price += 100
	}

	return price
}

func formatReceipt(movieName string, count int, priceTicket int) (string, int) {
	total := count * priceTicket
	receipt := fmt.Sprintf("Фильм: %s\nВсего билетов: %d\nЦена за один билет: %d\n", movieName, count, priceTicket)
	return receipt, total
}
