package main

import (
	"fmt"
)

// Doctor - declare struct
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// Nurse - uppercase properties means it can be exported
type Nurse struct {
	Age  int
	Name string
}

func main() {
	// instantiate doctor
	doctor1 := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	// can instantiate without using keys (not advised - hard to maintain )
	doctor2 := Doctor{
		1,
		"John Smith",
		[]string{
			"Joe Bloggs",
		},
	}

	nurse1 := Nurse{
		Age:  30,
		Name: "Dolores Price",
	}

	// anonymous struct - cant be used anywhere else
	aDoctor := struct{ name string }{name: "Jon Pertwee"}

	// this creates a copy not a reference
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"

	// this creates a reference to the original
	oneMoreDoctor := &anotherDoctor
	oneMoreDoctor.name = "Tom Jones"

	// print whole struct
	fmt.Println(doctor1)
	fmt.Println(doctor2)
	fmt.Println(nurse1)
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
	fmt.Println(oneMoreDoctor)

	// get property
	fmt.Println(doctor1.actorName)
	fmt.Println(doctor1.companions[0])
	fmt.Println(nurse1.Age)
	fmt.Println(aDoctor.name)
	fmt.Println(anotherDoctor.name)
	fmt.Println(oneMoreDoctor.name)
}
