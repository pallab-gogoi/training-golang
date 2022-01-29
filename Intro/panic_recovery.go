package main

import "fmt"

func main() {
	fmt.Println(divide(4, 0))
	//divide(2, 0)
}
func divide(a, b int) int {
	//defer func() { // this is a nameless or annonimous function
	//	fmt.Println(recover())

	//}()
	defer recoverB()
	if b == 0 {
		panic("Abey saale") //
	}

	return a / b
}

func recoverB() {
	if r := recover(); r != nil { // r is the value of panic
		fmt.Println("recovered", r)
	}
}

// if recover is not there panic will happen
// and termination of program will happen
// reflect.TypeOf() --> can be used to see data type
