package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter int = 0

// read-write mutex created below
// mutex - a lock that the application honors
// can protect code so that only one entity is able to manipulate that code/access data at one time
// if mutex is locked and something tries to alter/manipulate data, it must wait until its unlocked and it obtain the mutex
// RWMutex - unlimited entities can read data but only one can write. If anything is reading we cant write to it
var m = sync.RWMutex{}

// go run -race file.go - race flag can be used to detect race conditions
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		// data is accessed by many threads (counter)
		// functions do not run in sync unless mutex used
		// locks for to protecting against concurrent reading
		// this keeps code in sync
		// this has destroyed concurrency as its all running in sync :(
		m.RLock()
		go sayHello()
		// locks for writing
		m.Lock()
		go increment()
	}
	wg.Wait()

	// prints number of threads in program (-1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// number of threads can be altered (threads set to 1 below)
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// application can run faster with more threads (too many can incur mem overhead and slow scheduler)
	runtime.GOMAXPROCS(100)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	// unlocks single RLock
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	// unlocks mutex
	m.Unlock()
	wg.Done()
}
