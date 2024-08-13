package main

import (
	"fmt"
	"os"
)

// 	4. Запись строки в файл
//    • Описание: Напишите программу, которая записывает заданную строку в файл.
//    • Методы или функции:
//        ◦ func writeStringToFile(fileName, content string) error

func main() {
	err := writeStringToFile("fileWrite.txt", "Happy Happy Happy")
	if err != nil {
		fmt.Printf("writeStringToFile failed, err:%v\n", err)
	}

	fmt.Printf("writeStringToFile success\n")

}

func writeStringToFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(content)
	return nil

}
