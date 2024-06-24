package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	for _, char := range Itoa(Strlen(os.Args[1])) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func Strlen(arg string) (count int) {
	return len(arg)
}

func Itoa(n int) (str string) {
	var bytes []byte
	for x := n; x > 0; x /= 10 {
		bytes = append([]byte{byte(rune((n % 10) + 48))}, bytes...)
		n /= 10
	}
	str = string(bytes)
	return
}
