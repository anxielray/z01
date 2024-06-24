package main

import "fmt"

func Mirror(s string) string {
	var str string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			str += string('z' - (char - 'a'))
		} else if char >= 'A' && char <= 'Z' {
			str += string('Z' - (char - 'A'))
		}
	}
	return str
}

func main() {
	fmt.Print("input a string: ")
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println()
	fmt.Println(Mirror(s))
	fmt.Println()
	fmt.Println(`
	The string that you provided as your input
	 was converted by a magical mirror for each character.
	 This means that if your letter was 'a'  it shifts to 'z'
	 and if it were 'm' it becomes 'n'`)
}
