package main

import (
	"fmt"
	match "match/match"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		os.Stdout.WriteString("provide two strings to be compared")
		os.Exit(0)
	}
	s := os.Args[1]
	str := os.Args[2]
	fmt.Println(match.WdMatch(s, str))
}
