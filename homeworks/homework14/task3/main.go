package main

import (
	"bufio"
	"fmt"
	"os"
)

// 3. Счетчик слов в строке
//    • Описание: Напишите функцию, которая считает количество слов в строке.
//    • Методы или функции:
//        ◦ func countWords(s string) int

func main() {
	words := countWords("example.txt")
	fmt.Printf("%d words\n", words)

}

func countWords(s string) int {
	file, _ := os.Open(s)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0

	for scanner.Scan() {
		count++
	}
	return count
}
