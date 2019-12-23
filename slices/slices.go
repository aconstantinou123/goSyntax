package main

import (
	"fmt"
)

func main() {
	// slices dont have a fixed size
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	// slices are references types that point to an underlying array
	b := a
	b[0] = 0
	fmt.Println(a)

	//creating slices (can also be done on arrays)
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]   // slice of all elements
	e := c[3:]  // slice from 4th element to end
	f := c[:6]  // slice first 6 elements
	g := c[3:6] // slice the 4th, 5th and 6th elements
	// all arrays will change as all pointing to underlying array
	c[5] = 42
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// make function
	h := make([]int, 3, 100)
	fmt.Println(h)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h))

	i := []int{}
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))

	// append underlying array
	// all existing elements copied to new array with larger capacity
	// using make means you can start with a larger slice and only copy if necessary (expand capacity)
	i = append(i, 1)
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))
	// use spread operater to concat with another array
	i = append(i, []int{2, 3, 4, 5}...)
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))

	j := []int{1, 2, 3, 4, 5}
	// remove element from beginning
	k := j[1:]
	// remove element from end
	l := j[:len(j)-1]
	// remove element from the middle
	m := append(j[:2], j[3:]...)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	// unexpected behavior
	fmt.Println(j)
}
