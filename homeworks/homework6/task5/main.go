package main

import "fmt"

func main() {
	var currentCurrency float64 = 74.5
	rateUpdater(69.8, &currentCurrency)
	currencyConverter(1000, &currentCurrency)
	rateUpdater(70.8, &currentCurrency)
	currencyConverter(900, &currentCurrency)

}

// 5. Конвертация валют с ежемесячным обновлением курса
// • Начальный курс доллара к рублю равен 74.5.
// • Реализуйте функции для ежемесячного обновления курса валюты.
// • Реализуйте функцию для конвертации заданной суммы в рубли по текущему курсу.
// • Если курс становится ниже 70, выведите сообщение "Курс слишком низкий".

func rateUpdater(newCurrencyRate float64, currentCurrencyRate *float64) {
	fmt.Printf("Текущий курс равен %.2f\n", *currentCurrencyRate)
	*currentCurrencyRate = newCurrencyRate
	fmt.Printf("Курс изменен на %.2f\n", newCurrencyRate)
}

func currencyConverter(currencyToConvert float64, currentCurrencyRate *float64) {
	var convertedCurrency float64
	fmt.Printf("Вы хотите поменять %.2f долларов на рубли\n", currencyToConvert)
	if *currentCurrencyRate >= 70 {
		convertedCurrency = currencyToConvert * (*currentCurrencyRate)

		fmt.Printf("Вы получаете %.2f рублей\n", convertedCurrency)
	} else {
		fmt.Printf("Курс слишком низкий %.2f рублей\n", *currentCurrencyRate)
	}

}
