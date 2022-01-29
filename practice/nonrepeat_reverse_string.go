package main

import "fmt"

/*
problem 3
given a string, say s , print it in reverse manner eliminating
the repeated characterss and spaces

example 1:
Input: S = "GEEKS FOR GEEKS"
output: SKEGROF

example 2:
Input: S = "I AM AWESOME"
Output: "EMOSWAI"

*/
func reverse(s string) (reverse string) {
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}
func isContain(sample, char string) bool {
	if sample != "" {
		for _, value := range sample {
			if string(value) == char {
				return true
			}
		}
	}
	return false
}

func noRepeat(s string) (result string) {
	rvrs := reverse(s) // reverse string
	for _, value := range rvrs {
		s
		if isContain(result, string(value)) || string(value) == " " {
			continue
		} else {
			result += string(value)
		}

	}
	return
}

func main() {
	fmt.Println(noRepeat("GEEKS FOR GEEKS"))
}
