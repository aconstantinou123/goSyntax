package main

import (
	"fmt"
)

func main() {
	sayMessage("hello")

	for i := 0; i < 5; i++ {
		sayMessage2("Hello Go!", i)
	}

	greeting := "Hello"
	name := "Alex"
	sayGreeting(greeting, name)
	// name is unchanged
	fmt.Println(name)

	sayGreetingPointers(&greeting, &name)
	// name is changed
	fmt.Println(name)

	result := sum("hello", 1, 2, 3, 4, 5)
	fmt.Println(result)

	resultPointer := sumPointer(2, 3, 4, 5, 6)
	// can return a pointer from function (!) as variable is moved to heap after func data is garbage collected
	// deref pointer to get value
	fmt.Println(*resultPointer)

	greeting2 := namedReturn()
	fmt.Println(greeting2)

	// return err variable and exit if necessary
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// anonymous function
	func() {
		fmt.Println("Anonymous function")
	}()

	for i := 0; i < 5; i++ {
		// passing in variable rather than taking from outer scope prevents errors/strange behavior when dealing with async code
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	// assign anonymous function to variable
	// long declaration - var f func()
	f := func() {
		fmt.Println("Anonymous function as var")
	}
	f()

	// complex version
	var divide2 func(float64, float64) (float64, error)
	divide2 = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return a / b, nil
	}
	result3, err := divide2(10.0, 5.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result3)

	// func with callback
	callbackTest(f)
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayMessage2(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

// if types are the same, can write params like this
// data is copied if not a pointer
func sayGreeting(greeting, name string) {
	name = "Jim"
	fmt.Println(greeting, name)
}

// using pointers - a ref to the data is passed in
// much more efficient as not copying
func sayGreetingPointers(greeting, name *string) {
	*name = "John"
	fmt.Println(*greeting, *name)
}

// variatic functions - takes in variables and wraps up in slice
// can only have one variatic and it has to be last param
func sum(msg string, values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result, msg)
	return result
}

// returns address
func sumPointer(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// named return - instantiates return variable
func namedReturn() (greeting string) {
	greeting = "Hi"
	return
}

// multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// pass function in as callback
func callbackTest(callback func()) {
	callback()
}
