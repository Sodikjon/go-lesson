package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "Hello World Hello World Hello World Hello World"
	//wordsAndSymbolsCount(s)
	changeLetterToOther("b")
}

// 29. Подсчет слов и символов.
//     Напишите функцию, которая подсчитывает количество слов и символов в строке.

func wordsAndSymbolsCount(s string) {
	sSlice := strings.Split(s, " ")
	fmt.Printf("Количество слов: %d\n", len(sSlice))
	fmt.Printf("Количество символов %d\n", len(s))
}

// 28. Поиск всех анаграмм.
//	Напишите функцию, которая ищет все анаграммы строки в другой строке.

func findAnagrams(s string) {

}

// 26.  Генерация хеш-значения.
//	Напишите функцию, которая вычисляет хеш-значение строки с использованием алгоритма SHA-256.

func makeHash(s string) {
	//hash:= strings.
}

func changeLetterToOther(letter string) {
	fmt.Printf("Вы ввели букву %s\n", letter)
	b := letter[0]
	dc := int8(b) + 1
	l := byte(dc)
	fmt.Printf("Мы переделали букву %s\n", string(l))
}
