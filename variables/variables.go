package main

import (
	"fmt"
	"strconv"
)

// can declares vars up here but not using short syntax
var l float32 = 2.2

// group vars together
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 0
	season       int    = 11
)

var (
	counter int = 0
)

// uppercase vars are globalling scoped!
var I int = 42

// uppercase acronyms
var HTTPRequest string = "GET"

func main() {
	// declare
	var k int
	k = 22
	// long declaration
	var i float32 = 42
	// short declaration - go can guess the type
	j := 11.
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", l, l)

	fmt.Println(counter)
	var counter int = 1
	fmt.Println(counter)

	// coercing vars
	var n int = 1
	fmt.Printf("%v, %T\n", n, n)
	var o float32 = float32(n)
	fmt.Printf("%v, %T\n", o, o)
	// coercing strings
	var p int = 42
	fmt.Printf("%v, %T\n", p, p)
	var q string = strconv.Itoa(p)
	fmt.Printf("%v, %T\n", q, q)

}
