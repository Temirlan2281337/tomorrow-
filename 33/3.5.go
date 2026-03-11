package main

import "fmt"

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	var min uint
	if a < b {
		min = a
	} else {
		min = b
	}
	for i := min; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}

func main() {
	fmt.Println(Gcd(10, 42))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(10, 5))
	fmt.Println(Gcd(17, 3))
}
