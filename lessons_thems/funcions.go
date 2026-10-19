package main

import "fmt"

var totalFilms = 0 //Глобальная переменная

const maxCatalogSize = 100000

func addFilm() {
	totalFilms++ //Изменение локальной переменной из функции

}

func main() {
	addFilm()
	addFilm()
	fmt.Println(totalFilms)

	localMsg := "Only in main" //Локальная переменная
	fmt.Println(localMsg)

	outerVal := "внешняя переменная"
	fmt.Println(outerVal)
	//fmt.Println(innerVal) не видит, выдает ошибку
	{
		innerVal := "внутренняя переменная"
		fmt.Println(outerVal)
		fmt.Println(innerVal)
	}

	//Затенение
	price := 100
	fmt.Println(price)
	{
		price := 100
		fmt.Println(price) //Выдает внутреннюю переменную
	}
	fmt.Println(price) //Выдает внешнюю переменную

	price1 := 100

	fmt.Println(price1)
	{
		price1 = 200 //Изменили внешнюю переменную
		fmt.Println(price1)
	}
	fmt.Println(price1)
	addAllFilms(100)
	fmt.Println(addAllFilmsReturn(98))
	//Функция работает тооько с копией отправленной в неё переменной
	num := 500
	fmt.Println(returnModNum(num))
	fmt.Println(num)
}
func returnModNum(number int) int {
	number -= 10
	return number

}

// Передача и возврат параметров функции
// Передача аргумента в функцию
func addAllFilms(quantity int) {
	totalFilms += quantity
	fmt.Println(totalFilms)
}

// Передача и возврат
func addAllFilmsReturn(quantity int) int {
	totalFilms += quantity
	return totalFilms
}
