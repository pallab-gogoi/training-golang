// MAP : it's unordered
//
package main

import "fmt"

func main() {
	student_class := map[int]string{
		1: "E",
		2: "D",
		3: "C",
	}
	/*
		student_class2 := map[int]string{
			21: "G",
			22: "S",
			35: "T",
		}
	*/
	student_class[15] = "B"
	student_class[20] = "A"

	/* for roll_no, name := range student_class {
		fmt.Println("Roll no: ", roll_no, " Name: ", name)
	} */
	//delete.student_class[20]

	student, ok := student_class[15]
	if ok {
		fmt.Println(student)
	} else {
		fmt.Println("Key Invalid")
	}

}
