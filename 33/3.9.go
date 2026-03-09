package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		return // ничего не делаем, если аргументов не 3
	}

	str := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]

	if !strings.Contains(str, old) {
		fmt.Println(str)
		return
	}

	result := strings.ReplaceAll(str, old, new)
	fmt.Println(result)
}
