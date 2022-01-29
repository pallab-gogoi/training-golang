package main

import "fmt"

func main() {
	fmt.Println((fib(4)))
}
func fib(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	result := fib(number-1) + fib(number-2)
	return result
}

func fact(number int) int {
	if number == 1 || number == 0 {
		return 1
	}
	if number < 0 {
		fmt.Println("Invalid Number")
		return -1
	}
	result := number * fact(number-1)
	return result
}
