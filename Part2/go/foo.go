// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
	"sync"
)

// Control signals
const (
	GetNumber = iota
	Exit
)


//Workgroup
var wg sync.WaitGroup

func number_server(add_number <-chan int, control <-chan int, number chan<- int) {
	var i = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
		// TODO: receive different messages and handle them correctly
		// You will at least need to update the number and handle control signals.

		// New number channel
		case msg := <- add_number:
			i += msg // Increment/decrement i
			Println("Updated i to:", i)

		// Control Channel
		case msg := <- control:
			if msg == GetNumber {
				number <- i //Send value no number channel.
			}
			//		case msg := <- number: 
		default:

		}
	}
}

func incrementing(add_number chan<-int) {
	for j := 0; j<1000000; j++ {
		add_number <- 1
	}
	Println("Increment finished")
	wg.Done()
}

func decrementing(add_number chan<- int) {
	for j := 0; j<1000000; j++ {
		add_number <- -1
	}
	Println("Decrement finished")
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	// TODO: Construct the required channels
	addNumChan		:= make(chan int, 100) //Channel with 100 buffer size
	numChan			:= make(chan int)
	controlChan		:= make(chan int)
	// Think about whether the receptions of the number should be unbuffered, or buffered with a fixed queue size.

	// TODO: Spawn the required goroutines
	wg.Add(2)
	go number_server(addNumChan, controlChan, numChan)
	go incrementing(addNumChan)
	go decrementing(addNumChan)
	
	// TODO: block on finished from both "worker" goroutines
	wg.Wait() //Wait until workers done

	for len(addNumChan) != 0 { //Wait until channel is empty, probably not the prettiest way of doing it...
		
	}
	controlChan <- GetNumber
	Println("The magic number is:", <- numChan)
	controlChan <- Exit
	
}
