package main

import (
	"fmt"
	"sync"
)

// wait group synchronizes multiple go routines
var wg = sync.WaitGroup{}

// go run -race file.go - race flag can be used to detect race conditions
func main() {
	// a go routine is an abstraction over a green thread - we don't have to interact with the threads themselves
	// green threads have much less overhead that OS threads
	// go routines have small stack space and can be reallocated very quickly - cheap to create and destroy
	// the go keyword will spin up a 'green' thread and use it
	// scheduler assigns processing time on each thread created
	// main is not executing this function, the new thread is

	wg.Add(1)
	// telling wait group we are synchronizing to this go routine
	go sayHello()
	// tells wait group to wait for this function
	wg.Wait()

	msg := "Hello"
	wg.Add(1)
	go func() {
		// uses a closure to have access to variables in outer scope. Works even in a go routine. Not advised as it can create race conditions
		fmt.Println(msg)
		// decrements number of wait groups (back to zero)
		wg.Done()
	}()
	wg.Wait()

	// better to pass in variables - variable copied in when function invoked
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello")
	wg.Done()
	// decrements number of wait groups (back to zero)
}
