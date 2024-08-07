package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//str := "Hello, World!"
	//b := str[0]
	//fmt.Println(string(b))
	//
	//sub := str[7:12]
	//fmt.Println(sub)
	//
	//contains := strings.Contains(str, "World!!")
	//fmt.Println(contains)
	//index := strings.Index(str, "World")
	//fmt.Println(index)

	//s := "Hello, 世界"
	//bytes := []byte(s)
	//fmt.Println(bytes)
	//fmt.Println(string(bytes))
	//
	//for _, v := range s {
	//	fmt.Println(v)
	//	//fmt.Println(string(v))
	//
	//}

	// Пример строки с различными символами
	str := "Aé世😀"

	for len(str) > 0 {
		// Получаем первый байт
		r, size := utf8.DecodeRuneInString(str)

		fmt.Printf("Руна: %c, Unicode: U+%04X, Занимает %d байт\n", r, r, size)

		// Обрезаем строку, удаляя первый символ
		str = str[size:]
	}
}
