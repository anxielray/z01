package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	for _, char := range prevprime(os.Args[1]) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func findPrime(n int) (prime bool) {
	if n < 2 {
		return false
	}
	for x := 2; x < n; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}

func prevprime(arg string) (ans string) {
	var prev int
	prime := Atoi(arg)
	for i := prime; i >= prime-10; i-- {
		if !findPrime(i) {
			prev += 0
		} else {
			prev += i
			break
		}
	}
	ans = Itoa(prev)
	return
}

func Atoi(numstr string) (num int) {
	var numb, digit int = 1, 1
	for x := len(numstr) - 1; x >= 0; x-- {
		numb = int(rune(numstr[x]) - 48)
		num += (digit * numb)
		digit *= 10
		continue
	}
	return
}

func Itoa(num int) (str string) {
	var bytes []byte
	for x := num; x > 0; x /= 10 {
		bytes = append([]byte{byte(rune((num % 10) + 48))}, bytes...)
		num /= 10
	}
	str = string(bytes)
	return
}
