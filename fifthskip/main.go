package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	// fmt.Println(fifthAndSkip(os.Args[1]))
	for _, char := range fifthAndSkip(os.Args[1]) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func fifthAndSkip(s string) string {
	/*check if the string has enough length
	of arguments to use in the function.*/
	if len(s) < 5 {
		return "invalid input"
	}
	str := removeSpace(s)
	trimmed := ""
	count := 0
	for _, ch := range str {
		if count == 5 {
			trimmed += " "
			count = 0
			continue
		}
		trimmed += string(ch)
		count++
	}
	// return trimmed + "\n"
	return trimmed
}

func removeSpace(str string) string {
	result := ""
	for _, ch := range str {
		if ch != ' ' {
			result += string(ch)
		}
	}
	return result
}
