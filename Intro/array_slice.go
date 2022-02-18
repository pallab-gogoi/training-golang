package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th to the end
	d := a[:4]  //slice first 4 elements
	e := a[2:6] // slice the 3rd,4th,5th,6th elements
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	fmt.Println("d =", d)
	fmt.Println("e =", e)

	//fmt.Printf("type of b is %T", b)
}
