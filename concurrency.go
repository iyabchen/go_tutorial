package main

import (
	"fmt"
	"runtime"
	"sync"
	// "time"
)

func Go(channel chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	channel <- true

}

func Go1(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	channel := make(chan bool)

	go func() {
		fmt.Println("Go dog!")
		channel <- true
		close(channel)

	}()
	// time.Sleep(2 * time.Second) // better to use a constant to define a sleep length
	for v := range channel { // when doing iteration on channel, channel has to close to get it stop waiting
		fmt.Println(v)
	}
	//  <-channel // block main, until channel has something which was put by the child thread

	fmt.Println("Use channel with buffer ")
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}

	fmt.Println("Use wait group ")

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go1(&wg, i)
	}
	wg.Wait()

	fmt.Println("The usage of select ")

	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {

					o <- true

					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {

					o <- true

					break
				}
				fmt.Println("c2", v)

			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "world"
	close(c1)
	close(c2)

	<-o

}
