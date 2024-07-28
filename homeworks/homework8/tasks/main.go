package main

import "fmt"

func main() {
	arr1 := []int{5, -3, 10, 1, 8, -7, 2, 3, 10, -5}

	// 1. Найти максимальный элемент в массиве.
	//maxValue := arr1[0]
	//for _, v := range arr1 {
	//	if v > maxValue {
	//		maxValue = v
	//	}
	//}
	//fmt.Println("Максимум: ", maxValue)

	// 2. Найти минимальный элемент в массиве.
	//minValue := arr1[0]
	//for _, v := range arr1 {
	//	if v < minValue {
	//		minValue = v
	//	}
	//}
	//fmt.Println("Минимум: ", minValue)

	// 3. Подсчитать количество положительных чисел в массиве.
	//count := 0
	//for _, v := range arr1 {
	//	if v > 0 {
	//		count++
	//	}
	//}
	//fmt.Println("Количество положительных:", count)

	//// 4. Найти сумму всех элементов массива.
	//sum := 0
	//for _, v := range arr1 {
	//	sum += v
	//}
	//fmt.Println("Cумма элементов", sum)

	//// 5. Найти среднее значение всех элементов массива.
	//average := float64(sum) / float64(len(arr1))
	//fmt.Println("Среднее значение элементов", average)

	//// 6. Удалить все вхождения заданного числа из массива.
	//number1 := 10
	//result1 := []int{}
	//
	//fmt.Println("Исходный массив", arr1)
	//
	//for _, v := range arr1 {
	//	if v != number1 {
	//		result1 = append(result, v)
	//	}
	//}
	//fmt.Printf("Массив без числа %d %d", number1, result1)

	//// 7. Умножить все элементы массива на заданное число.
	//number2 := 5
	//result2 := make([]int, len(arr1))
	//
	//fmt.Println("Исходный массив", arr1)
	//
	//for i, v := range arr1 {
	//	result2[i] = v * number2
	//}
	//fmt.Printf("Результат умножения элементов массива на %d %d", number2, result2)

	// 8. Найти все индексы заданного числа в массиве.

	//// 9. Создать копию массива.
	//copiedArray := append([]int{}, arr1...)
	//fmt.Printf("Оригинальный массив: %d \nСкорированный массив: %d", arr1, copiedArray)

	// 10. Объединить два массива.
	arr1 = []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	unionizedArray := append(arr1, arr2...)
	fmt.Printf("Оригинальный массив: %d \nВторой массив: %d\nОбъедененный массив %d", arr1, arr2, unionizedArray)

}
