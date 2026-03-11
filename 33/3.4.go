package main

import "fmt"

func FishAndChips(n int) string {
	if n < 0 {
		return "number is negative"
	}
	if n%3 == 0 && n%2 == 0 {
		return "fish and chips"
	}
	if n%2 == 0 {
		return "fish"
	}
	if n%3 == 0 {
		return "chips"
	}

	return "non divisible"
}

func main() {
	fmt.Println(FishAndChips(11))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
