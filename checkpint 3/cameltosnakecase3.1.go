package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	res := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		if !(c >= 'a' && c <= 'z') && !(c >= 'A' && c <= 'Z') {
			return s
		}

		if i > 0 && c >= 'A' && c <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}

		if i == len(s)-1 && c >= 'A' && c <= 'Z' {
			return s
		}

		if i > 0 && c >= 'A' && c <= 'Z' {
			res += "_"
		}

		res += string(c)
	}

	return res
}

// func CamelToSnakeCase(s string) string {
// 	res := ""

// 	for i, c := range s {

// 		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') {
// 			return s
// 		}

// 		if i > 0 && c >= 'A' && c <= 'Z' && rune(s[i-1]) >= 'A' && rune(s[i-1]) <= 'Z' {
// 			return s
// 		}

// 		if i == len(s)-1 && c >= 'A' && c <= 'Z' {
// 			return s
// 		}

// 		if i > 0 && c >= 'A' && c <= 'Z' {
// 			res += "_"
// 		}

// 		res += string(c)
// 	}

// 	return res
// }

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
