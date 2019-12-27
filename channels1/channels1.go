package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// channels are normally always used in the context of go routines
// used to sync data between multiple go routines
func main() {
	// have to use make
	// channels are strongly typed
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		// 10 go routines spawned in total - all can use 1 channel
		wg.Add(2)
		// receiving go routine
		go func() {
			// receives data from channel
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		// sending go routine
		go func() {
			i := 42
			// sends data to channel
			// go routine is paused until channel can receive data - can create deadlocks if channel unavailable
			// only 1 message can be in the channel
			ch <- i
			// a copy is passed so manipulating data after doesn't matter
			i = 27
			wg.Done()
		}()
	}

	// Another example
	// both go routines below acting as readers and writers
	// not advised
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()
	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Add(2)
	//make channel receive only
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	// make channel send only
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)

	// buffered channels
	// second argument tells go to create a channel that can store up to 50 ints - both messages will be printed out
	// buffered channels are for when receivers and senders operate at different frequencies
	bufferedCh := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(bufferedCh)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		// close channel - for range above detects close. This avoids a deadlock
		// have to be careful - cannot send message after channel closed
		// cant detect that channel has been closed before sending message
		close(ch)
		wg.Done()
	}(bufferedCh)

	bufferedCh2 := make(chan int, 50)
	wg.Add(2)
	// manual detect of closed channel
	go func(ch <-chan int) {
		for {
			// ok can be used to detect when channel is closed
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(bufferedCh2)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(bufferedCh2)

	wg.Wait()
}
