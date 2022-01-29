package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Enter size of your array: ")
	var size, k int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your array is: ", arr)
	sort.Ints(arr)
	fmt.Println(" the array after sorting becomes", arr)
	fmt.Print("Enter value of K :")
	fmt.Scanln(&k)
	fmt.Println("the kth smallest integer is ", arr[k-1])
}
