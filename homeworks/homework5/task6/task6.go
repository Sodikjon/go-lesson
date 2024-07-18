package main

import "fmt"

func main() {
	// Напишите программу, которая выводит числа от 1 до 100, заменяя числа, кратные 3, на "Fizz",
	// числа, кратные 5, на "Buzz", а числа, кратные 3 и 5, на "FizzBuzz".

	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%s ", "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Printf("%s ", "Fizz")
		} else if i%5 == 0 {
			fmt.Printf("%s ", "Buzz")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}
