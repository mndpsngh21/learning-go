package basics

import (
	"fmt"
	"sync"
)

// var logChannel = make(chan int)
var intChannel = make(chan int, 50)
var cwg = sync.WaitGroup{}

func ChannelOps() {
	cwg.Add(2)
	go consumer(intChannel)
	go publisher(intChannel)
	cwg.Wait()

}

func publisher(ch chan<- int) {

	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)
	cwg.Done()
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println("Received data from channel ::", val)
	}
	cwg.Done()

}
