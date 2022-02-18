package main

import (
	"fmt"
	"time"
)

func main() {
	helloCh := make(chan string, 1)
	goodByeCh := make(chan string, 1)
	quitCh := make(chan bool)

	go receiveMessage(helloCh, goodByeCh, quitCh)

	go sendMessage(helloCh, "Hello world")

	time.Sleep(time.Second)

	go sendMessage(goodByeCh, "Good Bye world")
	result := <-quitCh
	fmt.Println("Message from quitCh=", result)
}
func sendMessage(ch chan<- string, message string) {
	ch <- message
}
func receiveMessage(helloCh, goodByeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case message := <-helloCh:
			fmt.Println("Message from helloCh: ", message)
		case message := <-goodByeCh:
			fmt.Println("Message from goodByeCh:  ", message)
		case <-time.After(time.Second * 2):
			fmt.Println("Nothing received in 2 secinds :   Eiting the receiver")
			quitCh <- true
			break
		}
	}
}
