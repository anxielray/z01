package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(0)
	}
	lcm := LCM(os.Args[1], os.Args[2])
	for _, char := range lcm {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func LCM(arg, arg2 string) (lcm string) {
	var (
		n   = Atoi(arg)
		n2  = Atoi(arg2)
		gcd = 1
	)
	var ans int
	for x := 1; x <= n && x <= n2; x++ {
		if n%x == 0 && n2%x == 0 {
			gcd = x
		}
		ans = (n * n2) / gcd
		lcm = Itoa(ans)
	}
	return
}

func Atoi(arg string) (num int) {
	var (
		number = 1
		digit  = 1
	)
	for x := len(arg) - 1; x >= 0; x-- {
		number = int(rune(arg[x]) - 48)
		num += (number * digit)
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
	str += string(bytes)
	return
}
