package main

import (
	"bytes"
	"fmt"
	"io"
)

// go has implicit interface implementation
func main() {
	// instantiate Writer as a concrete type - ConsoleWriter
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello, Go!"))

	// int counter with incrementer interface
	myInt := IntCounter(0)
	// if any methods uses a pointer receiver, interface must be implemented with a pointer not a value
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()

	// type conversion
	// casting to pointer to BufferedWriterCloser
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// failed conversion but avoiding panic
	// ok is bool that shows if conversion was successful
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	// empty interfaces - defined on the fly with no methods
	// everything can be cast into an object that has no methods on it (even primitives) - good for things that are type incompatible
	// must cast to another type to actually use it for anything useful
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("This is another test"))
		wc.Close()
	}

	// type switch - checks data type
	// commonly used with empty interface
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}

}

// Writer - declare interface
// naming convention (if single method) is to name interface the same as method + er
type Writer interface {
	// method signature
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// implementing interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// IntCounter - dont need to use structs to implement interfaces
// this uses an int type
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// more compicated example involving 2 interfaes (using writer interface defined above)
type Closer interface {
	Close() error
}

// WriterCloser - composed of 2 interfaces (same as embedding)
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// write data to console in series of 8 bytes
// pointer receiver
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// flushes buffer of final bytes
// pointer receiver
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// creates a pointer to a BufferedWriterCloser
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
