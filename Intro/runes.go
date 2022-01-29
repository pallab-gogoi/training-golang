package main

import (
	"fmt"
	"reflect"
)

func main() {
	var u rune
	u = 'a' // rune returns the ASCII value of "a"

	fmt.Println(u)
	fmt.Println(reflect.TypeOf(u))
	fmt.Println(string(u))
}
