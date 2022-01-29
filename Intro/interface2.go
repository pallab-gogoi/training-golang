package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Type= %T, value = %v \n", i, i)
	var num int
	var ok bool
	num, ok = i.(int)
	fmt.Println(num)
	fmt.Println(ok)
}
func main() {
	describe("Hello")
	describe(12)
}
