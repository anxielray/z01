package main

import (
	// "fmt"

	"github.com/01-edu/z01"
)

func main() {
	// fmt.Println("aBcDeFgHiJkLmNoPqRsTuVwXyZ")
	s := "aBcDeFgHiJkLmNoPqRsTuVwXyZ"
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
