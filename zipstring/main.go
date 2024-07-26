package main

import (
	"fmt"
	"os"
	"strconv"
)

func ZipString(s string) string {
	var str string
	n := len(s)

	if n == 0 {
		return ""
	}
	count := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			str += (strconv.Itoa(count))
			str += string(s[i-1])
			count = 1
		}
	}
	str += strconv.Itoa(count)
	str += string(s[n-1])
	return str
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		os.Exit(0)
	}
	fmt.Println(ZipString(os.Args[1]))
}
