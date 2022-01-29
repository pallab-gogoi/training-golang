package main

import (
	"fmt"
	"reflect"
)

//main function will be automatically called when you start the application
func main() {
	// var Varaiable_name datatype
	// var data int
	// data = 10

	// data1 := 100 // declare(automatic) the variable and assign value to it
	//	data := 1000 //redeclare the data variable
	// data = 1000
	// Data := 1000
	//	data1 = false // wrong assignment value type
	// fmt.Println(data)
	// fmt.Println(data1)
	// fmt.Println(Data)
	// data type
	//1. Primitave
	// int, float64, string, bool, complex, uint,  byte
	//2. Non-Primitive
	// struct, map, chan, interface, array, slice, rune, reflect, pointer
	//functional programming
	//recent feature update of java - from java 11
	// name := "rahul"
	// fmt.Println(name)

	// const pi = 3.14 //implici t typing
	const pi float32 = 3.14 //explicit typing
	const l int = 100
	// l = 10
	fmt.Println(pi)
	fmt.Println(l)
	fmt.Println(reflect.TypeOf(l))
}

// reflect.TypeOf() --> can be used to see data type
