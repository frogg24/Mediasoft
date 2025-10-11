package main

var m = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func convertRomanNumerals(s string) int {
	//преобразование в руны и проверка на корректность воода
	runeArray := []rune(s)
	for i := range runeArray {
		if _, exists := m[runeArray[i]]; !exists {
			return -1
		}
	}
	result := 0

	//формирование массива с интами на основе массива с рунами
	values := make([]int, 0, len(runeArray))
	for _, key := range runeArray {
		if value, exists := m[key]; exists {
			values = append(values, value)
		}
	}

	//проход по всем числами кроме последнего
	for i := 0; i < len(values); i++ {
		// если текущее число меньше следующего то приьавляю к результату их разность
		if i < len(values)-1 && values[i] < values[i+1] {
			result += values[i+1] - values[i]
			i++
		} else {
			result += values[i]
		}
	}
	return result
}
