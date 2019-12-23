package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	g.changeName("Alex")
	fmt.Println(g.name)
}

type greeter struct {
	greeting string
	name     string
}

// giving func the context on greeter struct so it can access it properties
// value receiver - struct is copied here - cant change data
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// pointer receiver - changes struct object
func (g *greeter) changeName(name string) {
	g.name = name
}
