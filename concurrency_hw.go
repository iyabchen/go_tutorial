package main

import (
	"fmt"
	"runtime"
	// "sync"
	// "time"
)

var c chan string

func Pingpong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From pingpong: hi, #%d", i)
		i++
	}
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	channel := make(chan string)

	go func() {

		fmt.Println("waiting for msg from main...")
		msg_send_child := "hello from child"
		channel <- msg_send_child

		msg_rec_child := <-channel
		fmt.Println(msg_rec_child)
		channel <- "msg received"
	}()

	// listen
	fmt.Println("waiting for msg from child...")
	msg_rec_main := <-channel
	fmt.Println(msg_rec_main)

	msg_send_main := "hello from main"
	channel <- msg_send_main

	<-channel

	c = make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From main: Hello, #%d", i)
		fmt.Println(<-c)
	}

}
