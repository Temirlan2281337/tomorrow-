package main

import (
	"unicode"
)

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	for i, r := range s {

		// запрещены цифры и символы
		if !unicode.IsLetter(r) {
			return s
		}

		// нельзя две заглавные подряд
		if i > 0 && unicode.IsUpper(r) && unicode.IsUpper(rune(s[i-1])) {
			return s
		}

		// нельзя заканчивать на заглавную
		if i == len(s)-1 && unicode.IsUpper(r) {
			return s
		}
	}

	result := ""

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result += "_"
		}
		result += string(r)
	}

	return result
}
