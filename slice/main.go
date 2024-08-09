// Instructions
// The function receives a slice of strings and one or more integers, and returns a slice of strings. The returned slice is part of the received one but cut from the position indicated in the first int, until the position indicated by the second int.
// In case there only exists one int, the resulting slice begins in the position indicated by the int and ends at the end of the received slice.
// The integers can be negative.
package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))

	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	result := []string{}
	length := len(a)
	if len(nbrs) == 1 {
		nb := nbrs[0]
		if nb < 0 {
			nb += length
		}
		result = append(result, a[nb:]...)
	} else if len(nbrs) == 2 {
		nb1 := nbrs[0]
		nb2 := nbrs[1]
		if nb1 < 0 || nb2 < 0 {
			nb1 += length
			nb2 += length
		} else if nb2 < nb1 {
			return result
		}
		result = append(result, a[nb1:nb2]...)
	}
	return result
}
