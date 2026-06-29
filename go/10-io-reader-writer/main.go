package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("hello, go!")
	counter := &countingReader{reader: input}
	upper := &upperWriter{writer: os.Stdout}
	io.Copy(upper, counter) // le do CountingReader, escreve no UpperWriter
	fmt.Printf("\n%d bytes lidos\n", counter.BytesRead())
}
