package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune(10)
		return
	}
	str := (PrintHex(os.Args[1])) //checkLength(PrintHex(os.Args[1]))
	for _, char := range str {
		// fmt.Printf("%x ", char)
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
	// println()
}

func PrintHex(s string) string {
	var str string
	for _, char := range s {

		str += numToHex(int(char)) + " " //runeToInt(char))
	}
	return str
}

// func runeToInt(r rune) int {
// 	fmt.Println(int(r - 48))
// 	return int(r - 48)
// }

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

// func checkLength(arg string) string {
// 	if len(arg) == 12 {
// 		z01.PrintRune(10)
// 	}
// 	return arg
// }
