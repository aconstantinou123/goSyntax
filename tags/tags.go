package main

import (
	"fmt"
	"reflect"
)

// Animal - name tagged (just a string, no value)
type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
