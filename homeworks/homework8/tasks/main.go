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

	//// 10. Объединить два массива.
	//arr1 = []int{1, 2, 3}
	//arr2 := []int{4, 5, 6}
	//unionizedArray := append(arr1, arr2...)
	//fmt.Printf("Оригинальный массив: %d \nВторой массив: %d\nОбъедененный массив %d", arr1, arr2, unionizedArray)

	//// 11. Поменять местами максимальный и минимальный элементы массива.
	//maxIndex, minIndex := 0, 0
	//
	//fmt.Println("Исходный массив:", arr1)
	//for i := 1; i < len(arr1); i++ {
	//	if arr1[i] > arr1[maxIndex] {
	//		maxIndex = i
	//	}
	//	if arr1[i] < arr1[minIndex] {
	//		minIndex = i
	//	}
	//}
	//arr1[maxIndex], arr1[minIndex] = arr1[minIndex], arr1[maxIndex]
	//fmt.Printf("Измененный массив: %d \n", arr1)

	//// 12. Проверить, является ли массив палиндромом.
	//arr1 = []int{-5, -3, 10, 2, 8, 2, 10, -3, -5}
	//isPalindrome := true
	//for i := 0; i < len(arr1)/2; i++ {
	//	if arr1[i] != arr1[len(arr1)-1-i] {
	//		isPalindrome = false
	//		fmt.Println("Не является палиндромом.")
	//		break
	//	}
	//}
	//if isPalindrome {
	//	fmt.Println("Уррра нашли палиндром.")
	//}

	//// 13. Найти второе наибольшее число в массиве.
	//firstMax, secondMax := arr1[0], arr1[0]
	//for _, v := range arr1[1:] {
	//	if v > firstMax {
	//		secondMax = firstMax
	//		firstMax = v
	//	} else if v > secondMax && v != firstMax {
	//		secondMax = v
	//	}
	//}
	//
	//fmt.Printf("Второе максимальное найдено: %d", secondMax)

	//// 14. Перевернуть массив.
	//fmt.Println("Исходный массив:", arr1)
	//for i := 0; i < len(arr1)/2; i++ {
	//	arr1[i], arr1[len(arr1)-1-i] = arr1[len(arr1)-1-i], arr1[i]
	//}
	//fmt.Printf("Измененный массив: %d \n", arr1)

	//// 15. Удалить дубликаты из массива.
	//result := []int{arr1[0]}
	//fmt.Println("Исходный массив:", arr1)
	//for i := 1; i < len(arr1); i++ {
	//	isDuplicate := false
	//	for j := 0; j < len(result); j++ {
	//		if arr1[i] == result[j] {
	//			isDuplicate = true
	//			break
	//		}
	//	}
	//	if !isDuplicate {
	//		result = append(result, arr1[i])
	//	}
	//}
	//fmt.Printf("Результат: %d \n", result)

	// 16. Переместить все нули в конце массива, сохраняя порядок ненулевых элементов.
	arr1 = []int{5, -3, 0, 1, 8, -7, 2, 3, 0, -5}
	fmt.Println("Исходный массив:", arr1)
	nonZeroIndex := 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != 0 {
			arr1[nonZeroIndex], arr1[i] = arr1[i], arr1[nonZeroIndex]
			nonZeroIndex++
		}
	}
	fmt.Printf("Измененный массив: %d \n", arr1)

	// 17. Найти пересечение двух массивов.
	// 18. Проверить, является ли массив подмножеством другого массива.
	// 19. Объединить два отсортированных массива в один отсортированный.
	// 20. Найти длину самого длинного подмассива, в котором все элементы различны.
}
