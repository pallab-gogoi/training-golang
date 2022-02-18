package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	counter := 200
	wait.Add(counter)
	defer wait.Wait()
	for i := 0; i < counter; i++ {
		go hello(&wait, i)
	}
}
func hello(wait *sync.WaitGroup, i int) {
	fmt.Println(i)
	wait.Done()
}
