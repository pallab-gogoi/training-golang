package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	go func(ch chan string) {
		mess := <-ch
		ch <- "Hey from Anonymous"
		fmt.Println(mess)
		fmt.Println(1)
	}(channel)
	message := "Hey from Main func"
	channel <- message
	time.Sleep(time.Second * 5)
	message = <-channel
	fmt.Println(message)
}
