package basics

import (
	"fmt"
	"sync"
)

func Routines() {
	waitGroupLogic()
	mutexLogic()
}

// for block code we need locking so we use wait group for wait execution
var wg = sync.WaitGroup{}

func waitGroupLogic() {
	wg.Add(1)       // tell waitgroup that there is one task to be executed
	go randomCode() // make any function goroutine execution by just adding go keyword before it
	wg.Wait()       // wait for all goroutines to be completed

}

func randomCode() {
	fmt.Println("Executing code")
	wg.Done() // this tells wait group that task is completed so decrement worker counter
}

/*
*
Mutex logic:

	Mutex helps to define flow, they work as blocker for any code, it is similar locking functionality

*
*/
var m = sync.RWMutex{}

func mutexLogic() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		m.Lock() // adding lock here as otherwise order can not be determined
		go sayHello(i)
	}
	wg.Wait()
}

func sayHello(n int) {

	fmt.Println("Hello sir ::", n)
	m.Unlock()
	wg.Done() // this tells wait group that task is completed so decrement worker counter
}
