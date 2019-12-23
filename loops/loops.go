package main

import (
	"fmt"
)

func main() {
	// basic loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// multiple variables
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// i declared outside of loop
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	// incrementer not needed in loop syntax
	for i < 10 {
		i++
		fmt.Println(i)
	}

	//infinite for loop
	for {
		fmt.Println(i)
		i++
		if i < 20 {
			fmt.Println("less than 20")
			continue
		}
		if i == 20 {
			fmt.Println("i is 20")
			break
		}
	}

	// Label - allows block to be labelled. Break then breaks out of both loops
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// for range loop
	s := []int{1, 2, 3}
	for k, v := range s {
		// index and values
		fmt.Println(k, v)
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	word := "hello, go!"
	for k, v := range word {
		fmt.Println(k, string(v))
	}

	// ignore variable
	for _, v := range word {
		fmt.Println(string(v))
	}

}
