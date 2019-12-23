package main

import (
	"fmt"
	"log"
)

func main() {
	// recover allows programming to keep going if it encounters a panic
	fmt.Println("start")
	panicker()
	// functions higher up call stack can keep executing
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		// recover function
		if err := recover(); err != nil {
			log.Println("Error", err)
			// can rethrow panic
			// panic(err)
		}
	}()
	panic("something bad happened")
	// function recovers but stops execution
	fmt.Println("done panicking")
}
