package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	for _, char := range Revwstr(os.Args[1]) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func Revwstr(arg string) (rev string) {
	var last int
	var first rune = ' '
	for i, char := range arg {
		if i == 0 && char != first {
			arg = string(first) + arg
		}
	}
	for x := len(arg) - 1; x >= 0; x-- {
		if string(arg[x]) == " " {
			last = x
			for i := last + 1; i < len(arg); i++ {
				rev += string(arg[i])
			}
			rev += " "
			arg = arg[:last]
			rev += Revwstr(arg)
			break
		}
	}
	return
}
