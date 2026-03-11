package main

import "fmt"

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}

	if n < 0 {
		n = n * -1
	}

	count := 0

	for n > 0 {
		n /= base
		count++
	}

	num := string(n)

	newNum := int(num[:])
	print(newNum)
	// if num[0] == '-'

	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
