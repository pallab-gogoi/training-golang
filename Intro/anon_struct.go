/*
anonymous structure do not have a name
these are created only once
*/
package main

import "fmt"

func main() {
	s := struct {
		name   string
		age    int
		height float32
	}{}
	s.age = 21
	s.name = "john"
	s.height = 6.4
	fmt.Printf("%v is %v years old and height is %v", s.name, s.age, s.height)
}
