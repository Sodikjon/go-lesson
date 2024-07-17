package main

import "fmt"

func main() {
	PrintGreeting()
}

func PrintGreeting() {
	var dayType string
	fmt.Println("Введите время дня ")
	fmt.Scanf("%s", &dayType)
	switch dayType {
	case "день":
		fmt.Printf("Добрый %s", dayType)

	case "вечер":
		fmt.Printf("Добрый %s", dayType)
	}
}
