package main

import (
	"fmt"
)

func main() {
	fmt.Println(addfront("Hello, ", []string{"world"}))
}

/*
Write a function that takes a string and a slice of strings, this function will return a new slice of string with the given string prepended

If the given string is empty you only need to return the given slice
*/
func addfront(s string, sta []string) []string {
	if s == "" {
		return sta
	}
	return append([]string{s}, sta...)
}
