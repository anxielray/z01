package main

import (
	"github.com/01-edu/z01"
)

func add(m, n int) int {
	return m + n
}
func multiply(m, n int) int {
	return m * n
}
func divide(m, n int) int {
	return m / n
}
func main() {
	a := []int{500, 2}
	for _, char := range reduceInt(a, add) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
	for _, char := range reduceInt(a, multiply) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
	for _, char := range reduceInt(a, divide) {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func reduceInt(a []int, f func(int, int) int) (ans string) {
	result := a[0]
	for _, num := range a[1:] {
		result = f(result, num)
	}
	ans = Itoa(result)
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
