package main

import "fmt"

const (
	SUB_PRICE = 100
	SUB_NAME  = "Premium"
)

func main() {
	discount := 15.0
	months := 12
	priceMonths := SUB_PRICE - ((SUB_PRICE / 100) * discount)
	total := priceMonths * float64(months)

	fmt.Println("======GoFlix======")
	fmt.Printf("Тариф: %s\nБазовая цена: %d рублей в месяц\n", SUB_NAME, SUB_PRICE)
	fmt.Printf("Месяцев: %d\n", months)
	fmt.Printf("Скидка: %.0f%% \nИтого: %.2f рублей\n", discount, total)

}
