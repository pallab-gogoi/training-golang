/*
A pointer to the struct can be created with the & operator or the new keyword.
pointer is dereferenced with * operator
*/
package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	p := Point{3, 4}
	pp := &p // pp is a pointer to the struct Point
	pp.x = 1
	pp.y = 2
	fmt.Println(p)
}
