package main

import "fmt"

func main() {
	// 2. Проверка наличия ключа
	//Описание: Создайте map с несколькими элементами и напишите функцию, которая проверяет наличие заданного ключа.

	m := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m)

	fmt.Println(hasKey(4, m))

}

func hasKey(key int, m map[int]string) bool {
	if _, ok := m[key]; ok {
		return true
	} else {
		return false
	}
}
