package main

import (
	"bufio"
	"fmt"
	"os"
)

// 5. Чтение первой строки файла
//    • Описание: Напишите программу, которая читает первую строку из текстового файла.
//    • Методы или функции:
//        ◦ func readFirstLine(fileName string) (string, error)

func main() {
	firstLine, err := readFirstLine("example.txt")
	if err != nil {
		fmt.Printf("Error reading first line: %v\n", err)
	}

	fmt.Printf("%s\n", firstLine)
}

func readFirstLine(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}
