package lesson2

// FizzBuzz возвращает FizzBuzz, Fizz, Buzz или пустую строку для обычного числа.
// Для n <= 0 возвращает invalid.
func FizzBuzz(n int) string {
	// TODO: Реализовать через switch без выражения.
	switch {
	case n <= 0:
		return "invalid"
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	}
	return ""
}

// Sum возвращает сумму. Для пустого среза и nil возвращает 0.
func Sum(values []int64) int64 {
	// TODO: Пройти по значениям циклом.
	// abc
	if values == nil || len(values) == 0 {
		return 0
	}
	var sum int64
	for _, v := range values {
		sum += v
	}
	return sum
}

// Min возвращает минимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Min(values []int64) int64 {
	// TODO: Учесть пустой ввод, затем сравнить элементы.
	if values == nil || len(values) == 0 {
		return 0
	}
	minimum := values[0]
	for _, v := range values {
		if v < minimum {
			minimum = v
		}
	}
	return minimum
}

// Max возвращает максимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Max(values []int64) int64 {
	// TODO: Учесть пустой ввод, затем сравнить элементы.
	if values == nil || len(values) == 0 {
		return 0
	}
	maximum := values[0]
	for _, v := range values {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}

// Average возвращает среднее как float64. Для пустого среза и nil возвращает 0.
func Average(values []int64) float64 {
	// TODO: Выполнить преобразования до деления.
	if values == nil || len(values) == 0 {
		return 0
	}
	var sum float64
	for _, v := range values {
		sum += float64(v)
	}
	return sum / float64(len(values))
}

// Category возвращает invalid, small, regular или large по сумме в копейках.
func Category(amount int64) string {
	// TODO: Проверить границы через switch без выражения.
	switch {
	case amount < 0:
		return "invalid"
	case amount <= 9999:
		return "small"
	case amount <= 49999:
		return "regular"
	}
	return "large"
}

// RuneCount считает кодовые точки Unicode в строке через range.
func RuneCount(text string) int {
	// TODO: Не использовать len(text) как количество кодовых точек.
	var count int
	for _ = range text {
		count++
	}
	return count
}
