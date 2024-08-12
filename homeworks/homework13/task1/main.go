package main

import (
	"fmt"
	"strings"
)

//   1.  Длина строки и Количество слов
//    • Описание: Реализуйте интерфейс StringProcessor, который будет содержать методы Length() и WordCount().
//   Реализуйте структуру MyString, которая будет работать с текстом и реализует этот интерфейс.
//    • Методы:
//        ◦ Length() int для получения длины строки.
//        ◦ WordCount() int для подсчета количества слов

type StringProcessor interface {
	Length() int
	WordCount() int
}

type MyString struct {
	text string
}

func (ms MyString) Length() int {
	return len(ms.text)
}

func (ms MyString) WordCount() int {
	words := strings.Fields(ms.text)
	return len(words)
}

func main() {
	text := MyString{"Это пример текста для подсчета слов и длины."}

	var processor StringProcessor = text

	fmt.Printf("Длина строки: %d\n", processor.Length())
	fmt.Printf("Количество слов: %d\n", processor.WordCount())
}
