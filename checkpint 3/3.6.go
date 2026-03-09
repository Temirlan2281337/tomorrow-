package main

import "fmt"

func HashCode(dec string) string {
	n := len(dec)
	result := ""

	for i := 0; i < n; i++ {
		c := (int(dec[i]) + n) % 127
		if c < 33 {
			c += 33
		}
		result += string(rune(c))
	}

	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
