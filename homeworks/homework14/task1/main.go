package main

import (
	"fmt"
	"os"
)

// 1. Подсчет символов в файле
//    • Описание: Напишите программу, которая считывает файл и подсчитывает количество символов в нём.
//    • Методы или функции:
//        ◦ func countCharacters(fileName string) (int, error)

func main() {
	characters, err := countCharacters("example.txt")
	if err != nil {
		fmt.Printf("Error counting characters: %s\n", err)
	}
	fmt.Printf("Characters: %d\n", characters)
}

func countCharacters(fileName string) (int, error) {
	// Открытие файла для чтения
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, err
	}
	defer file.Close()

	// Чтение данных из файла
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, err
	}
	return n, nil
}
