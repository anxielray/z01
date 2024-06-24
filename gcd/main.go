package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(0)
	}
	answer := GCD(os.Args[1], os.Args[2])
	for _, char := range answer {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func GCD(arg, arg2 string) (gcd string) {
	var (
		n   = Atoi(arg)
		n2  = Atoi(arg2)
		ans = 1
	)
	for x := 1; x <= n && x <= n2; x++ {
		if n%x == 0 && n2%x == 0 {
			ans = x
		}
	}
	gcd = Itoa(ans)
	return
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
