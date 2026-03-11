package main

import (
	"fmt"
)

// DigitLen считает количество цифр в числе n для системы счисления base
func DigitLen(n, base int) int {
	// 1. Проверка корректности базы (от 2 до 36)
	if base < 2 || base > 36 {
		return -1
	}

	// 2. Обработка нуля: в числе "0" всегда одна цифра
	if n == 0 {
		return 1
	}

	// 3. Убираем минус, так как он не влияет на количество разрядов
	if n < 0 {
		n = -n
	}

	count := 0

	// 4. Основной цикл: делим число на базу, пока оно не станет 0
	// Каждое деление — это одна "откушенная" цифра
	for n > 0 {
		n /= base
		count++
	}

	return count
}

func main() {
	// Проверяем разные сценарии
	fmt.Println(DigitLen(100, 10))  // Выведет: 3
	fmt.Println(DigitLen(100, 2))   // Выведет: 7 (1100100)
	fmt.Println(DigitLen(-100, 16)) // Выведет: 2 (64)
	fmt.Println(DigitLen(100, 1))   // Выведет: -1
}

// $ go run . | cat -e
// 3$
// 7$
// 2$
// -1$
