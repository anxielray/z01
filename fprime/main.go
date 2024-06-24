package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(Fprime(num))
}

func Fprime(num int) (factors []int) {
	var prefactors []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			prefactors = append(prefactors, i)
		}
	}
	for _, facts := range prefactors {
		if findPrime(facts) {
			factors = append(factors, facts)
		}
	}
	return
}

func findPrime(n int) bool {
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
