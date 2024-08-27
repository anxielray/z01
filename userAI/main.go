package main

import (
	"fmt"
	"os"
)

func main() {
	AskNameAge()
}

func AskNameAge() {

	var name string
	var age int
	var emoji string = string(rune(128522))

	fmt.Print("Hello! My name is Anxiel, What's yours: \n")
	fmt.Scanln(&name)

	fmt.Println("\nGreat,! Wat a nice and cool name you got! What is your age, If you don't mind ? ")
	fmt.Scanln(&age)
	if age <= 0 || age >= 130 {
		fmt.Println("Oops! You seem to have given an invalid age number!")
		os.Exit(0)
	}
	
	fmt.Printf("I hope am getting this right %s !,Your name is  %s!\n\nAnd you are %d years of age.\n", emoji, name, age)
}
