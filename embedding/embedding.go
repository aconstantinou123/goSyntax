package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

// Bird - embeds an animal (no inheritance in go)
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// one way of instantiating
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)

	// literal syntax - need to declare using internal struct
	a := Bird{
		Animal:   Animal{Name: "Eagle", Origin: "America"},
		SpeedKPH: 50,
		CanFly:   true,
	}

	fmt.Println(a)
	fmt.Println(a.CanFly)
}
