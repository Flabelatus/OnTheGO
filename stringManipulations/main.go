package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "alpha alpha alpha alpha alpha alpha"

	str = replaceNth(str, "alpha", "beta", 4)
	fmt.Println(str)

	// newString := "Go is a great programming language. Go for it!"

	// if strings.Contains(newString, "Go") {
	// 	newString = strings.Replace(newString, "Go", "Golang", 1)
	// }

	// fmt.Println(newString)

	// fmt.Println()

	// // string comparison
	// if "Alpha" > "Absolute" {
	// 	fmt.Println("Alpha is greater than Absolute")

	// } else {
	// 	fmt.Println("Alpha is not greater than Absolute")
	// }

	// badEmail := " me@here.com  "
	// badEmail = strings.TrimSpace(badEmail)

	// // to check and see that now it does not contain white space
	// fmt.Printf("=%s=", badEmail)
	// fmt.Println()

}

func replaceNth(s, old, new string, n int) string {
	// index
	i := 0

	for j := 1; j <= n; j++ {
		// x is the index of the searched string in the strings slice
		x := strings.Index(s[i:], old)
		if x < 0 {
			// the string did not exist in the strings (s)
			break
		}

		i = i + x
		// every step in the loop the value of j increaments until it is equal to the
		// requested index value at n
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}

		i += len(old)

	}

	return s
}
