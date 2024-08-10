// write a program that prints the last word in an argument eg. "hello world" world
package main

import (
	//"strings"
	"github.com/01-edu/z01"
)

func main() {
	s := "Anxiel is the god of nature called Pan in foklores"
	for _, cha := range FirstWord(s) {
		z01.PrintRune(cha)
	}
	z01.PrintRune(10)
}

func FirstWord(s string) string {
	// stringArray := strings.Fields(s)
	// return stringArray[0]
	var count int = 1
	var str string
	for i, char := range s {
		if char == ' ' {
			count++
			if count == 2 {
				str += s[0:i]
			}
		}
	}
	return str
}
