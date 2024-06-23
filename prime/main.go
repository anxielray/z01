package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	fmt.Println(prime(Atoi(os.Args[1])))
}

func Atoi(arg string) (num int) {
	var digit, number = 1, 0
	for x := len(arg) - 1; x >= 0; x-- {
		number = int(rune(arg[x]) - 48)
		num += (digit * number)
		digit *= 10
	}
	return
}

func prime(n int) bool {
	if n < 2 {
		return false
	}
	for x := 2; x <= int(math.Sqrt(float64(n))); x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}
