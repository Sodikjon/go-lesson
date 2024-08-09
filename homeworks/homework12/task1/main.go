package main

import "fmt"

func main() {
	// 1. Создание и вывод map
	//	Описание: Создайте map для хранения строк и их длин, добавьте несколько элементов и выведите содержимое.
	stringsMap := make(map[string]int)
	s1 := "Hello World"
	s2 := "Hi Go"
	s3 := "Hello HumoLab"

	stringsMap[s1] = len(s1)
	stringsMap[s2] = len(s2)
	stringsMap[s3] = len(s3)

	fmt.Println(stringsMap)
}
