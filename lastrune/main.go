package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	z01.PrintRune(rune(os.Args[1][len(os.Args[1])-1]))
}
