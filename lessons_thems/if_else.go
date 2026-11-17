package main

import "fmt"

func main() {
	var ticketsLeft int = -2
	fmt.Println("Жопа", ticketsLeft)
	if ticketsLeft > 0 {
		fmt.Println("Билеты есть")
	} else if ticketsLeft == 0 {
		fmt.Println("Билетов нет")
	} else {
		fmt.Println("Система сломалась")
	}
}
