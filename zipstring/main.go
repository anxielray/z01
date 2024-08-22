package main

import (
	"fmt"
	"os"
	"strconv"
)

func ZipString(s string) string {
	var str string
	n := len(s) // stores the length of the string s...

	// if no argment is passed, the program returns an empty string...
	if n == 0 {
		return ""
	}

	// initialise count to 1...
	count := 1
	for i := 1; i < n; i++ { // set an index that runs from the index 1 to the last index...
		if s[i] == s[i-1] {
			count++
		} else { // if the prefixing character is not equal to the previous one, count is retained as 1.
			str += (strconv.Itoa(count))
			str += string(s[i-1])
			count = 1 // the value of the count is reset back to 1... and the loop moves to the next i value...w
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
