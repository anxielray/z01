package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	for _, char := range RepeatAlpha(os.Args[1]) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func RepeatAlpha(alph string) (repeat string) {
	var count int
	for _, char := range alph {
		if char >= 'a' && char <= 'z' {
			count = int((char - 'a') + 1)
			for i := 1; i <= count; i++ {
				repeat += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			count = int((char - 'A') + 1)
			for i := 1; i <= count; i++ {
				repeat += string(char)
			}
		} else {
			repeat += string(char)
		}
	}
	return
}
