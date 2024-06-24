package main

import (
	"fmt"
	"os"
)

func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	fmt.Println(Compare(os.Args[1], os.Args[2]))
}
