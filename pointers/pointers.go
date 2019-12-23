package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	// use deref to reassign value
	*b = 27
	fmt.Println(a)
	// get address of a
	fmt.Println(&a)
	// dereference b
	fmt.Println(*b)

	// go does not normally allow pointer arithmetic
	// can be done with unsafe package
	d := [3]int{1, 2, 3}
	e := &d[0]
	f := &d[1]
	fmt.Printf("%v %p %p\n", d, e, f)

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	// use new keyword - values initialized to zero
	var ms2 *myStruct
	// pointer that is not initalized automatically initialized to nil
	fmt.Println(ms2)
	ms2 = new(myStruct)
	fmt.Println(ms2)

	// must deref to get at and change fields
	(*ms2).foo = 3
	fmt.Println((*ms2).foo)

	// deref syntax is optional
	ms2.foo = 5
	fmt.Println(ms2.foo)

	// sharing slices always shares pointers to the underlying array
	// both slices change value
	g := []int{1, 2, 3}
	h := g
	fmt.Println(g, h)
	g[1] = 42
	fmt.Println(g, h)

	// sharing arrays makes a copy
	// only i changes value
	i := [3]int{1, 2, 3}
	j := i
	fmt.Println(i, j)
	i[1] = 42
	fmt.Println(i, j)

	// sharing maps also only shares a pointer to the data
	// both maps change the value of foo
	k := map[string]string{"foo": "bar", "baz": "buz"}
	l := k
	fmt.Println(k, l)
	k["foo"] = "qux"
	fmt.Println(k, l)
}

type myStruct struct {
	foo int
}
