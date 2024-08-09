package main

import "fmt"

func main() {
	// 30. Реализация map для хранения срезов строк
	//Описание: Реализуйте map, где ключи – это строки, а значения – срезы строк. Добавьте функции для добавления, удаления и получения срезов.

	slicesMap := make(map[string][]string)
	slicesMap["greetings"] = []string{"Hello", "Hi", "Whats up"}

	fmt.Println("Initial map of slices", slicesMap)
	fmt.Println()

	byesSlice := []string{"Good bye", "Bye", "GB"}

	addToMap("byes", byesSlice, slicesMap)
	fmt.Println(slicesMap)
	fmt.Println()

	getMapItems(slicesMap)
	fmt.Println()

	fmt.Println()
	deleteFromMap("greetings", slicesMap)
	fmt.Println(slicesMap)

}

func addToMap(key string, slice []string, m map[string][]string) {
	m[key] = slice

	fmt.Printf("key:%s with value slice:%v was added to map \n", key, slice)
}

func deleteFromMap(key string, m map[string][]string) {
	if _, ok := m[key]; ok {
		delete(m, key)
		fmt.Printf("key:%s was deleted from map\n", key)
	} else {
		fmt.Printf("key:%s does not exist\n", key)
	}
}

func getMapItems(m map[string][]string) {
	fmt.Println("Here is my map")
	for k, v := range m {
		fmt.Printf("key:%s, slice:%v\n", k, v)
	}
}
