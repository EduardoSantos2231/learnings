package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"time"
)

func main() {

	var w io.Writer // (type=nil, value=nil) → w == nil: true
	fmt.Printf("Type of w: %T\n", w)
	var buf *bytes.Buffer // nil pointer
	fmt.Printf("\nU must receive an erro because is not possible to dereference a nil val\n")
	time.Sleep(time.Second * 2)
	w = buf // (type=*bytes.Buffer, value=nil) → w != nil: true!
	fmt.Printf("\n%#X", *buf)
	// w.Write([]byte("hello")) // panic!
}

func safeWrite(w io.Writer, data []byte) (int, error) {
	var invalidWritterImplementation = errors.New("invalid implementation of io.Writer")
	if w == nil {
		return 0, invalidWritterImplementation
	}
	val := reflect.ValueOf(w)
	kind := val.Kind()
	switch kind {
	case reflect.Chan, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Slice:
		if val.IsNil() {
			return 0, invalidWritterImplementation
		}
	}
	return w.Write(data)
}
