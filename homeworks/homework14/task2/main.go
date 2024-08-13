package main

import (
	"bufio"
	"fmt"
	"os"
)

// 2. Подсчет строк в файле
//    • Описание: Напишите программу, которая считает количество строк в текстовом файле.
//    • Методы или функции:
//        ◦ func countLines(fileName string) (int, error)

func main() {
	lines, err := countLines("example.txt")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("lines: %d\n", lines)
}

func countLines(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 0

	for scanner.Scan() {
		count++
	}
	return count, nil
}
