package main

import (
	"fmt"
)

func main() {
	// ... - make array of size of elements passed in. Can also be a number
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #2: %v\n", students[1])
	fmt.Printf("Number of Students: %v\n", len(students))

	//Array of arrays
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	// reassigning and array copies data and makes new array
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// creating a pointer to the original array
	c := &a
	c[1] = 6
	fmt.Println(a)
	fmt.Println(c)
}
