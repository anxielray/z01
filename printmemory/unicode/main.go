package main

import (
	"github.com/01-edu/z01"
	"os"
	"unicode"
)

func printMemory(arr [10]byte) {
	base := "0123456789abcdef"
	for i := 0; i < len(arr); i++ {
		div := int(arr[i]) / len(base)
		mod := int(arr[i]) % len(base)
		z01.PrintRune(rune(base[div]))
		z01.PrintRune(rune(base[mod]))
		if i != 3 && i != 7 && i != 9 {
			z01.PrintRune(' ')
		}
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune(10)
		}
	}
	for i := 0; i < len(arr); i++ {
		if unicode.IsGraphic(rune(arr[i])) {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
		z01.PrintRune(10)
	}
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	arr := [10]byte(os.Args[1])
	str := printMemory(arr)
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
