package main

import "fmt"

func main() {
	// Напишите программу, которая выводит таблицу умножения на 3 от 1 до 10.
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", i, 3, i*3)
	}
}
