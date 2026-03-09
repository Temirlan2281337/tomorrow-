package main

import "fmt"

func LastWord(s string) string {
	i := len(s) - 1

	// пропускаем пробелы в конце
	for i >= 0 && s[i] == ' ' {
		i--
	}

	end := i

	// ищем начало слова
	for i >= 0 && s[i] != ' ' {
		i--
	}

	return s[i+1:end+1] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
