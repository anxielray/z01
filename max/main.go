package main

import "github.com/01-edu/z01"

func main() {
	answer := Max([]int{9, 2, 90, 12, 43, 0, 1, 50})
	for _, char := range answer {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func Max(a []int) (maximum string) {
	var max int
	max = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	maximum = Itoa(max)
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
