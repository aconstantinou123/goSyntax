package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

// signal only channel - common pattern
// struct with no fields - no mem is allocated
// can only show that a message was sent or received
var doneCh = make(chan struct{})

func main() {
	go logger()
	// could use defer to gracefully shutdown channel after main finished
	// defer func() {
	// 	close(logCh)
	// }()

	// data passed to channel (log entry instance)
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	// syntax for defining and initializing struct on the fly
	// done channel called at end to close channel
	doneCh <- struct{}{}
}

// logger function monitors app for messages in channel and prints out
// logging logic can be kept in 1 place
func logger() {
	for {
		select {
		// log channel prints messages
		case entry := <-logCh:
			// entries in channel are logEntry struct
			fmt.Printf(
				"%v - [%v] %v\n",
				entry.time.Format("2006-01-02T15:04:05"), entry.severity,
				entry.message,
			)
		// done channel breaks loop
		case <-doneCh:
			break

		}
		// adding a default case will make statement non-blocking
	}
}
