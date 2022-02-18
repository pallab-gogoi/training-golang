/*
struct is a user-defined data
related data
can be compared to a light-weight class without inheritance
defined with the struct keyword

syntax-
type name struct{}
*/
package main

import "fmt"

func main() {
	u := User{
		age:    24,
		name:   "john doe",
		height: 6.4,
	}

	fmt.Printf("%v is %v years old with height %v \n", u.name, u.age, u.height)
}

type User struct {
	name   string
	age    int
	height float32
}
