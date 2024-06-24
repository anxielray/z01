package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	for _, char := range Rostring(os.Args[1]) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func Rostring(arg string) (rotate string) {
	var first int
	for i, char := range arg {
		if char == ' ' {
			first = i
			break
		}

	}
	rotate += arg[(first + 1):]
	rotate += " "
	rotate += arg[:first]
	return
}
