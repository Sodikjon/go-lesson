package main

import (
	"fmt"
)

func main() {
	// Напишите программу, которая выводит пирамиду из символов '*' высотой n, введённого пользователем.

	var height int

	fmt.Print("Введите высоту пирамиды: ")
	fmt.Scan(&height)

	printPyramid(height)
}

func printPyramid(height int) {
	for i := 1; i <= height; i++ {
		// Печатаем пробелы перед звездочками
		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}

		// Печатаем звездочки
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		// Переходим на новую строку
		fmt.Println()
	}
}
