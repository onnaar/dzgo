package dz

import "unicode/utf8"

// PairInput соответствует объекту nums/target из условия контеста.
type PairInput struct {
	Nums   []int
	Target int
}

// longestPrefix возвращает самый длинный общий префикс строк.
func longestPrefix(words []string) string {
	// TODO: реализовать поиск префикса по кодовым точкам Unicode.
	var lPrefix = ""
	if len(words) == 0 {
		return lPrefix
	}
	//var lPrefix string = ""
	//lPrefix := ""
	//var lPrefix string
	word := words[0]
	for pos, letter := range word {
		for _, w := range words[1:] {
			if pos >= len(w) {
				return lPrefix
			}
			r, _ := utf8.DecodeRuneInString(w[pos:])
			if letter != r {
				return lPrefix
			}
		}
		lPrefix += string(letter)
	}
	return lPrefix
}

// pairIndices возвращает индексы двух разных элементов по возрастанию.
func pairIndices(input PairInput) [2]int {
	// TODO: найти единственную пару с суммой input.Target.
	valMap := make(map[int]int)
	for i, num := range input.Nums {
		j, ok := valMap[input.Target-num]
		if ok {
			return [2]int{j, i}
		}
		valMap[num] = i
	}
	return [2]int{}
}

// hasDuplicate сообщает, встречается ли какое-либо значение повторно.
func hasDuplicate(nums []int) bool {
	// TODO: проверить наличие повторов.
	seen := make(map[int]int)
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = num
	}
	return false
}

// compactSorted создаёт независимый слайс уникальных значений.
func compactSorted(nums []int) []int {
	seen := make(map[int]int)
	var res []int
	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			res = append(res, num)
			seen[num] = num
		}
	}
	// TODO: сохранить порядок и не изменять исходный слайс.
	return res
}

// balancedBrackets проверяет пары и порядок закрытия скобок.
func balancedBrackets(text string) bool {
	// TODO: обработать строку из символов ()[]{}.
	var stack []rune
	for _, r := range text {
		switch r {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != r {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
