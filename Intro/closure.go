package main

import "fmt"

func main() {
	variable := 0 // declaring a variable

	count := func() int { // anonymous function
		variable += 1
		return variable
	}

	fmt.Println(count())
}
