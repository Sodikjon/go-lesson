package main

import "fmt"

func main() {
	// Напишите программу, которая проверяет, является ли число простым.

	var a int
	fmt.Println("Введите цисло: ")
	fmt.Scan(&a)

	d := 2
	for a%d != 0 {
		d += 1
	}
	if a == d {
		fmt.Println("Целое число")
	} else {
		fmt.Println("Цисло не целое")
	}
}
