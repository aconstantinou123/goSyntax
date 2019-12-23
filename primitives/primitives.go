package main

import (
	"fmt"
)

func main() {
	// booleans
	var m bool
	var n bool = true
	y := 1 == 1
	z := 2 == 1
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)

	// numeric types
	o := 42
	var p int8 = 55
	var q int16 = 55
	var r int32 = 55
	var s int64 = 2323525235234
	var t uint16 = 22
	fmt.Printf("%v, %T\n", o, o)
	fmt.Printf("%v, %T\n", p, p)
	fmt.Printf("%v, %T\n", q, q)
	fmt.Printf("%v, %T\n", r, r)
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", t, t)

	// arithmetic operations - must match data type
	fmt.Println(p + int8(q))

	// logic operators - using binary
	a := 10            // 1010
	b := 3             // 0011
	fmt.Println(a & b) // 0010 - matching numbers
	fmt.Println(a | b) // 1011 - either number has digit
	fmt.Println(a ^ b) // 1001 - digit in one but not other
	fmt.Println(a | b) // 0100 - digit in neither

	// bit shifting
	c := 8              // 1000
	fmt.Println(c << 3) // 1000000
	fmt.Println(c >> 3) // 1

	// floating point numbers - defaults to float32
	d := 2.2
	e := 21.7e32
	f := 26.7E32
	var g float32 = 5.4345
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", g, g)

	// complex type
	var h complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", h, h)

	// text type - utf8
	i := "this is a string"
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", string(i[2]), string(i[2]))
	// concat a string
	j := "Hello "
	k := "everyone"
	fmt.Printf("%v, %T\n", j+k, j+k)
	// string to bytes
	l := []byte(i)
	fmt.Printf("%v, %T\n", l, l)
	// rune - utf32
	aa := 'a'                      // rune is int32
	fmt.Printf("%v, %T\n", aa, aa) // int32 97
}
