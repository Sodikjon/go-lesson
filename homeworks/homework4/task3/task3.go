package main

import "fmt"

func main() {
	PrintTrafficLight()
}

func PrintTrafficLight() {
	var lightColor string
	fmt.Println("Введите цвет ")
	fmt.Scanf("%s", &lightColor)
	switch lightColor {
	case "красный":
		fmt.Println("Стоп")
	case "желтый":
		fmt.Printf("Внимание")
	case "зеленый":
		fmt.Printf("Идите")
	}
}
