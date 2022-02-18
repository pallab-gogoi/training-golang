// structure fields can be functions
package main

import "fmt"

func main() {
	u := User{
		name:       "John Doe",
		age:        21,
		occupation: "smugler",
		info: func(name string, occupation string, age int) string {
			return fmt.Sprintf("%s is %d years old and he is a %s\n", name, age, occupation)
		},
	}
	fmt.Printf(u.info(u.name, u.occupation, u.age))
}

type Info func(string, string, int) string
type User struct {
	name       string
	occupation string
	age        int
	info       Info
}
