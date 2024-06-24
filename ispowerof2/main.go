package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num := Atoi(os.Args[1])
	for _, char := range isPowerOfTwo(num) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func isPowerOfTwo(n int) string {
	var str string
	if n <= 0 {
		str = "false"
	}
	if ((n & (n - 1)) == 0) == true {
		str = "true"
	} else {
		str = "false"
	}
	return str
}

func Atoi(numstr string) (num int) {
	var numb int
	var digit int = 1
	for x := len(numstr) - 1; x >= 0; x-- {
		numb = int(rune(numstr[x]) - 48)
		num += (digit * numb)
		digit *= 10
		continue
	}
	return
}

func Itoa(num int) (str string) {
	numstr := []byte{}
	for x := num; x > 0; x /= 10 {
		numstr = append([]byte{byte(rune((num % 10) + 48))}, numstr...)
		num /= 10
	}
	str += string(numstr)
	return
}
