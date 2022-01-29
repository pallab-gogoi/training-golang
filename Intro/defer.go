package main

import "fmt"

func main() {
	defer f() // last defer  runs first
	defer x() // f should execute as the last statement of the main function

	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)

}

func f() {
	fmt.Println("Test test")
}
func x() {
	fmt.Println("Last")
}
