package main

// 1. Вариант 1
//   - Напишите функцию, которая принимает срез целых чисел и возвращает новый срез,
//     содержащий только уникальные элементы из исходного среза.
//   - Реализуйте функцию, которая принимает срез строк и возвращает срез,
//     в котором строки упорядочены по длине. Если длины строк совпадают, то строки должны упорядочиваться
//     в лексикографическом порядке.
//   - Напишите функцию, которая принимает срез чисел и возвращает срез, где каждое число умножено на 2,
//     но при этом сумма всех чисел в исходном срезе должна оставаться постоянной.
func main() {

}

func uniqueElements(slice []int) []int {
	var uniqueElements []int
	for i := 0; i < len(slice); i++ {
		uniqueElements = append(uniqueElements, slice[0])
		for _, v := range slice {
			if slice[i] != v {
				uniqueElements = append(uniqueElements, v)
			}
		}
	}
	return uniqueElements
}

func wordsAsc(strings []string) []string {
	ascStrings := make([]string, len(strings))
	for i := 0; i < len(strings); i++ {
		for _, s := range strings {
			if len(strings[i]) <= len(s) {
				ascStrings = append(ascStrings, s)
			}
		}
	}
	return ascStrings
}

func doubledNums(s1 []int) []int {
	newSlice := make([]int, 0)
	for _, v := range s1 {
		v = v * 2
		newSlice = append(newSlice, v)
	}
	return newSlice
}
