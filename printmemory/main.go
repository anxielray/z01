package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 { // ensure we have the correct number of arguments in the command line... Print a new line!
		z01.PrintRune(10)
		return
	}
	str := (PrintHex(os.Args[1]))
	for _, char := range str {
		// fmt.Printf("%x ", char) // Short fmt. way to  print the address of a character...
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
	// println()
}

func PrintHex(s string) string {
	var str string
	var count int
	for _, char := range s {
		str += numToHex(int(char)) + " " // concatenate each string(character) in the original string to a new proceeded by a new line...
		count++
		if count == 4 {
			str += "\n"
		}
	}
	return str
}

func numToHex(num int) (hex string) {
	hexMap := map[int]string{
		0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f",
	}
	for num > 0 {
		hex = hexMap[num%16] + hex
		num /= 16
	}
	return
}
