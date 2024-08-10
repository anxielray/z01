// write a program that prints the last word in an argument eg. "hello world" world
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	s := "Anxiel is the god of nature called Pan in foklores"
	for _, cha := range LastWord(s) {
		z01.PrintRune(cha)
	}
	z01.PrintRune(10)
}

func LastWord(s string) string {
	// stringArray := strings.Fields(s)
	// return stringArray[len(stringArray)-1]
	var count int = 1
	var focusIndex int
	for i, char := range s {
		if char == ' ' {
			count++
			focusIndex = i
		}
	}
	return s[focusIndex+1:]
}
