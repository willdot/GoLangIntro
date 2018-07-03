package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	//runInOrder()
	runConcurrent()
	//runParallel()

}

func runInOrder() {
	// This will run the first function, wait 5 seconds and then print hello, go to the second function and print will
	fmt.Println("Running in order")
	func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}() // these () at the end make them self executing

	func() {
		fmt.Println("Will")
	}()
}

func runConcurrent() {
	// This will run both concurrently
	// The wait group is created to say how many go routines will be used. Then the first function runs, until the sleep at which point the second function starts as it gets the thread.
	// Then when the second thread exits it reports that it's finished to the wait group and the first function gets the thread and will continue running once the sleep has finished.
	// Then when the first function has exits, it reports it's finished and that's it.
	// The waitgroup.wait() at the bottom is there to stop this function exiting before both functions have exited
	fmt.Println("Running concurrently")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}() // these () at the end make them self executing

	go func() {
		defer waitGroup.Done()

		fmt.Println("Will")
	}()

	waitGroup.Wait()
}

func runParallel() {
	// This will run both functions at the same time.
	fmt.Println("Running concurrently")
	runtime.GOMAXPROCS(2) // This ups the amount of threads available
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}() // these () at the end make them self executing

	go func() {
		defer waitGroup.Done()

		fmt.Println("Will")
	}()

	waitGroup.Wait()
}
