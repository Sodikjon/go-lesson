package main

import "fmt"

func main() {
	// 3. Обновление значения по ключу
	//Описание: Напишите функцию для обновления значения в map по заданному ключу.

	m := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m)

	updateByKey(2, "z", m)

	fmt.Println(m)
}
func updateByKey(key int, newValue string, m map[int]string) {
	if _, ok := m[key]; ok {
		m[key] = newValue
		fmt.Printf("key:%d updated by value:%s \n", key, m[key])
	} else {
		fmt.Printf("key:%d not exist", key)
	}
}
