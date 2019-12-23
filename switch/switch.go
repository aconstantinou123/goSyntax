package main

import (
	"fmt"
)

func main() {
	// can create a statement
	switch i := 2 + 3; i {
	// multiple cases
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// tagless syntax - cases can overlap
	i := 1
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		// falls through to the next case
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// type switch
	var j interface{} = 3
	switch j.(type) {
	case int:
		fmt.Println("i is an int")
		if j == 3 {
			fmt.Println("I hate 3")
			break
		}
		fmt.Println("Print as well")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string ")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("i is another type")
	}
}
