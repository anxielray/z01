package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	value, err := strconv.Atoi(os.Args[1])
	value2, _ := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("ERROR!")
		os.Exit(0)
	}
	if (int64(value) > 9223372036854775807 ||
		int64(value) < -9223372036854775807) ||
		(int64(value2) > 9223372036854775807 ||
			int64(value2) < -9223372036854775807) {
		fmt.Println("ERROR!")
	}
	operator := os.Args[2]
	fmt.Println(Doop(int64(value), operator, int64(value2)))
}

func Doop(value int64, operator string, value2 int64) (answer int64) {
	switch operator {
	case "+":
		answer = value + value2
	case "-":
		answer = value - value2
	case "*":
		answer = value * value2
	case "/":
		if value2 == 0 {
			fmt.Println("no division by 0...")
			os.Exit(0)
		}
		answer = value / value2
	case "%":
		if value2 == 0 {
			fmt.Println("no modulo by 0...")
			os.Exit(0)
		}
		answer = value % value2
	default:
		os.Exit(0)
	}
	if answer > 9223372036854775807 {
		fmt.Println("ERROR!")
		os.Exit(0)
	}
	return
}
