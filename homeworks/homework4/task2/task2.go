package main

import "fmt"

func main() {
	PrintWeather()
}

func PrintWeather() {
	var weatherType string
	fmt.Println("Введите погоду ")
	fmt.Scanf("%s", &weatherType)
	switch weatherType {
	case "солнечно":
		fmt.Println("Солнечно")
	case "облачно":
		fmt.Printf("Облачно")
	case "дождливо":
		fmt.Printf("Дождливо")
	}
}