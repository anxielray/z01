package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	answer := PrevPrimeSum(os.Args[1])
	for _, char := range answer {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
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

func FindPrime(n int) bool {
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

func PrevPrimeSum(arg string) (prevStr string) {
	var prev int
	prime := Atoi(arg)
	for i := prime; i >= 2; i-- {
		if !FindPrime(i) {
			prev += 0
		} else {
			prev += i
		}
	}
	prevStr = Itoa(prev)
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
