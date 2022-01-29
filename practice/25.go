package main

import "fmt"

func main() {

}
func a() (int, string) {
	return 21, "Pallab"
}
func b(args ...int) (bool, int) {
	for _, v := range args {
		fmt.Println(v)
	}
	return true, 11
}
func c(w int, name *string) (i int, j string) {
	i = 101
	j = "Pallab"
	w = 201
	*name = "gogoi"
	return
}
