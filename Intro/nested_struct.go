package main

import "fmt"

func main() {
	s := User{}
	s.name = "John Doe"
	s.age = 21
	s.address.country = "USA"
	s.address.city = "NY"
	s.address.zip = 1001

	fmt.Printf("%v is %v years old and address is %v,%v,%v", s.name, s.age, s.address.city, s.address.zip, s.address.country)
}

type Address struct {
	city    string
	zip     int
	country string
}
type User struct {
	name    string
	age     int
	address Address
}
