package main

import "fmt"

func main() {
	// Напишите программу, которая выводит первые 10 чисел последовательности Фибоначчи.

	fibonachi, f2 := 0, 1

	for i := 0; i < 10; i++ {
		fmt.Println(fibonachi)

		temp := fibonachi
		fibonachi = f2
		f2 = temp + f2
	}
}
