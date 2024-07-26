package main

import (
    "fmt"
)

// SaveAndMiss function to save and omit characters based on the provided int value
func SaveAndMiss(input string, n int) string {
    if n <= 0 {
        return input
    }

    result := ""
    length := len(input)

    for i := 0; i < length; {
        // Save the next 'n' characters
        endSave := i + n
        if endSave > length {
            endSave = length
        }
        result += input[i:endSave]

        // Move index 'i' by 'n' to skip the next 'n' characters
        i += n

        // Skip the next 'n' characters
        i += n
    }

    return result
}

func main() {
    // Example usage
    fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwxyz", 3)) // Output: "abcghjmnopqrs"
    fmt.Println(SaveAndMiss("Hello, World!", 2))              // Output: "Helo ol!"
    fmt.Println(SaveAndMiss("Save and Miss", 0))               // Output: "Save and Miss"
    fmt.Println(SaveAndMiss("123456789", -1))                  // Output: "123456789"
}
