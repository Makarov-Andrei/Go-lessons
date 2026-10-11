package main

import "fmt"

const (
	HOUR      = 3600
	BytesInGB = 1073741824.0
)

func main() {
	fmt.Println("====Привет, это конвертер====")
	var time int64 // Время в секундах
	var size int64 // Размер в байтах
	fmt.Println("Введите значения")
	fmt.Scan(&time, &size)
	hours := time / HOUR
	minutes := (time % HOUR) / 60
	seconds := time % 60
	var sizeFile = float64(size) / BytesInGB
	fmt.Printf("Фильм длиться: %d часов, %d минут и %d секунд.\n", hours, minutes, seconds)
	fmt.Printf("Твой файл весит: %.2f гигов.\n", sizeFile)
}
