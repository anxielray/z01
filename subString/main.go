package main

import "fmt"

func main() {
	s := "Iamthegreabstofall"
	fmt.Println(Contains(s))
}

//Create a function to chek for  the existence of subtring in a string passed as a parameter, without using the standard package
func Contains(s string) bool {
	sub := "abc" //declare the substring...
	subLen := len(sub)
	sLen := len(s)

	//base case...
	if subLen > sLen {
		return false
	}

	for i := 0; i <= sLen-subLen; i++ {
		if s[i:i+subLen] == sub {
			return true
		}
	}

	return false
}
