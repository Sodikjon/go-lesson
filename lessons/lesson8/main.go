package main

import (
	"fmt"
)

func main() {
	//var arr [5]int
	//arr2 := [3]string{"a", "b", "c"}
	//arr3 := [...]int{1, 2, 3, 4, 5}
	//
	//fmt.Println(arr)
	//
	//for i, v := range arr2 {
	//	fmt.Println(i, v)
	//}
	//
	//for i, v := range arr3 {
	//	fmt.Println(i, v)
	//}

	//slice1 := []int{1, 2, 3}
	//slice2 := []int{4, 5}
	//slice3 := append(slice1, slice2...) // добавление среза slice2 в срез slice1
	//slice4 := append(slice2, slice1...)
	//
	//fmt.Println(len(slice3), cap(slice3))
	//fmt.Println(len(slice4), cap(slice4))

	//arr := [5]int{1, 2, 3, 4, 5}
	//slice := arr[1:4]    // срез содержит элементы со 2-го по 4-й (индексы 1, 2, 3)
	//fmt.Println(slice)   // [2 3 4]
	//
	//slice2 := slice[1:2] // срез на основе другого среза
	//fmt.Println(slice2)  // [3]

	//src := []int{1, 2, 3}
	//dst := make([]int, len(src))
	//n := copy(dst, src) // копирование всех элементов из src в dst
	//fmt.Println(dst)    // [1 2 3]
	//fmt.Println(n)      // 3

	//arr := [5]int{1, 2, 3, 4, 5}
	//slice := arr[1:4]
	//slice[0] = 10
	//fmt.Println(arr) // [1 10 3 4 5]

	//var nilSlice []int
	//fmt.Println(nilSlice == nil)              // true
	//fmt.Println(len(nilSlice), cap(nilSlice)) // 0 0

	////Abdulhalim Ergashev, [26.07.2024 19:29]
	////// Создайте массив из 7 целых чисел и инициализируйте его значениями от 1 до 7.
	//arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	//fmt.Println(arr)

	////Abdulhalim Ergashev, [26.07.2024 19:29]
	////// Создайте массив из 5 строк и инициализируйте его значениями "a", "b", "c", "d", "e".
	//arr2 := [5]string{"a", "b", "c", "d", "e"}
	//fmt.Println(arr2)

	////Abdulhalim Ergashev, [26.07.2024 19:30]
	////// Создайте массив из 4 логических значений и инициализируйте его значениями true, false, true, false.
	//arr3 := [4]bool{true, false, true, false}
	//fmt.Println(arr3)

	////Abdulhalim Ergashev, [26.07.2024 19:34]
	////// Создайте массив из 10 целых чисел без инициализации и выведите его на экран.
	//var arr4 [10]int
	//fmt.Println(arr4)

	////Abdulhalim Ergashev, [26.07.2024 19:34]
	////// Создайте массив из 4 логических значений и выведите значения по индексам 1 и 3.
	//arr5 := [4]bool{true, false, true, false}
	//fmt.Println(arr5[1])
	//fmt.Println(arr5[3])

	//Abdulhalim Ergashev, [26.07.2024 19:34]
	//// Создайте массив из 3 логических значений и измените значение первого элемента на false.
	//arr6 := [3]bool{true, false, true}
	//fmt.Println(arr6)
	//arr6[0] = false
	//fmt.Println(arr6)

	////Abdulhalim Ergashev, [26.07.2024 19:35]
	////// Создайте массив из 4 строк и выведите все его элементы с помощью цикла for range.
	//arr7 := [4]string{"a", "b", "c", "d"}
	//
	//for i, v := range arr7 {
	//	fmt.Println(i, v)
	//}

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	idx := 0
	for i, v := range a {
		if a[0] < v && v < a[9] {
			idx = i
		}
	}
	fmt.Println(idx)

	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	idx1 := 0
	for i, v := range a {
		if a1[0] < v && v < a1[9] {
			idx1 = i
			break
		}
	}
	fmt.Println(idx1)
}
