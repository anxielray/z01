package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	for _, char := range Rot14(os.Args[1]) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func Rot14(arg string) (rot14 string) {
	for _, char := range arg {
		if (char >= 'a' && char <= 'l') || (char >= 'A' && char <= 'L') {
			rot14 += string(char + 14)
		} else if (char >= 'm' && char <= 'z') || (char >= 'M' && char <= 'Z') {
			rot14 += string(char - 12)
		} else {
			os.Exit(0)
		}
	}
	return
}
