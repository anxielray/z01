package main

import (
	"fmt"
)

// Split splits a string s into substrings separated by the specified delimiter sep.
// It returns a slice of the substrings.
func Split(s, sep string) []string {
	if sep == "" {
		return []string{s} // If separator is empty, return the original string in a slice
	}

	var result []string
	start := 0
	for {
		// Find the index of the next occurrence of sep
		index := indexOf(s[start:], sep)
		if index == -1 {
			// If sep is not found, add the remaining substring to the result
			result = append(result, s[start:])
			break
		}
		// Add the substring before sep to the result
		result = append(result, s[start:start+index])
		// Move the start index past the found sep
		start += index + len(sep)
	}

	return result
}

// indexOf returns the index of the first occurrence of sep in s, or -1 if not found.
func indexOf(s, sep string) int {
	if len(sep) == 0 {
		return -1 // Return -1 if the separator is empty
	}
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			return i // Return the index if sep is found
		}
	}
	return -1 // Return -1 if sep is not found
}

func main() {
	// Example usage of the custom Split function
	str := "hello,world,this,is,go"
	sep := ","
	result := Split(str, sep)

	fmt.Println("Original string:", str)
	fmt.Println("Separator:", sep)
	fmt.Println("Split result:", result)
}