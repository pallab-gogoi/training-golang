package main

import (
	"fmt"
)

// func main() {
// 	channel := make(chan int)

// 	go func() {
// 		channel <- 1
// 		time.Sleep(time.Second * 2)
// 		channel <- 2
// 		close(channel)
// 	}()
// 	for ch := range channel {
// 		fmt.Println(ch)
// 	}

// }

func sendInt(ch chan int) {
	ch <- 1
}
func sendBool(ch chan bool) {
	ch <- true
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	go sendBool(ch2)
	go sendInt(ch1)
	select {
	case getInt := <-ch1:
		fmt.Println(getInt)
	case getBool := <-ch2:
		fmt.Println(getBool)
	default:
		fmt.Println("default")
	}
}
