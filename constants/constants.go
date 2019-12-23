package main

import (
	"fmt"
)

const a int16 = 27

//enumerated constants
// iota works as a counter. Increments +1 everytime its used
const (
	// compiler works out pattern
	c = iota
	d
	e
)

const (
	f = iota
)

const (
	// value ignored (0)
	_ = iota + 5
	catSpecialist
	dogSpecialist
)

const (
	_  = iota // ignore first value
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// using bitshifting for access roles
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	const x string = "hello"
	fmt.Printf("%v %T\n", x, x)

	// can overwrite (shadow) constants that are set at package level
	fmt.Printf("%v %T\n", a, a)
	const a int32 = 42
	fmt.Printf("%v %T\n", a, a)

	// cant set constants to something that is calculated at runtime e.g
	// const a = float64 = math.Sin(1.57)

	// inferred type
	const b = 32
	fmt.Printf("%v %T\n", b, b) // 32 int
	// type can be changed by compiler of not declared
	fmt.Printf("%v %T\n", b+a, b+a) // 74 int32

	fmt.Printf("%v %T\n", c, c) // 0
	fmt.Printf("%v %T\n", d, d) // 1
	fmt.Printf("%v %T\n", e, e) // 2
	// iota scoped to constant block
	fmt.Printf("%v %T\n", f, f) // 0

	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)
	fmt.Printf("%v\n", catSpecialist)
	fmt.Printf("%v\n", dogSpecialist)

	filesize := 4000000000.
	fmt.Printf("%0.2fGB\n", filesize/GB)
	fmt.Printf("%v\n", KB)
	fmt.Printf("%v\n", MB)
	fmt.Printf("%v\n", GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is Headquarters? %v\n", isAdmin&roles == isHeadquarters)
}
