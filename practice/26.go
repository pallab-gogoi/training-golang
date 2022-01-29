// MAP : it's unordered
package main

import "fmt"

func main() {
	student_class := map[int]string{
		1: "E",
		2: "D",
		3: "C",
	}
	student_class[15] = "B"
	student_class[20] = "A"

	/* for roll_no, name := range student_class {
		fmt.Println("Roll no: ", roll_no, " Name: ", name)
	} */

	student, ok := student_class[135]
	if ok {
		fmt.Println(student)
	} else {
		fmt.Println("Key Invalid")
	}
}
